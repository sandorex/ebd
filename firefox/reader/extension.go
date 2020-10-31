// Copyright 2020 Aleksandar Radivojevic
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package reader

import (
	"encoding/json"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/sandorex/ebd/common"
	"github.com/sandorex/ebd/errors"
	"time"
)

// Extension represents Firefox extension
type Extension struct {
	ID          string
	Name        string
	Version     string
	Enabled     bool
	Description string
	AddonURL    string
	InstallDate time.Time

	// These are exclusive to Firefox
	DisabledByUser bool
	Author         string
}

// ParseExtensions parses extensions from the file
func ParseExtensions(data []byte) (map[string]Extension, error) {
	var v map[string]*json.RawMessage
	err := json.Unmarshal(data, &v)
	if err != nil {
		return nil, err
	}

	version, err := common.ReadVersionFromJSON("schemaVersion", v)
	if err != nil {
		return nil, err
	}

	fn, ok := extensionVersions[version]
	if !ok {
		return nil, errors.ErrUnsupportedSchemaVersion{Version: version}
	}

	return fn(v)
}

// TODO get extensions.json version 31 and make the parse function for 32
var extensionVersions = map[int](func(data map[string]*json.RawMessage) (map[string]Extension, error)){
	31: parseExtensionsV31,
}

func parseExtensionsV31(data map[string]*json.RawMessage) (map[string]Extension, error) {
	rawIdentities, ok := data["addons"]
	if !ok {
		return nil, errors.ErrJSONMissingKeys([]string{"addons"})
	}

	var v []struct {
		ID             *string `json:"id" valid:"required"`
		Version        *string `json:"version" valid:"required"`
		Type           *string `json:"type" valid:"required"`
		Enabled        *bool   `json:"active" valid:"required"`
		DisabledByUser *bool   `json:"userDisabled" valid:"required"`
		Hidden         *bool   `json:"hidden" valid:"required"`
		Location       *string `json:"location" valid:"required"`
		DateInstalled  *int64  `json:"installDate" valid:"required"`
		DefaultLocale  *struct {
			Name        *string `json:"name" valid:"required"`
			Description *string `json:"description" valid:"required"`
			Creator     *string `json:"creator" valid:"required"`
		} `json:"defaultLocale" valid:"required"`
	}

	err := json.Unmarshal(*rawIdentities, &v)
	if err != nil {
		return nil, err
	}

	// convert to common type
	extensions := map[string]Extension{}
	for _, i := range v {
		_, err := govalidator.ValidateStruct(i)
		if err != nil {
			return nil, err
		}

		// skip non-extension addons, builtin addons, hidden addons
		if *i.Type != "extension" || *i.Location != "app-profile" || *i.Hidden {
			continue
		}

		extensions[*i.ID] = Extension{
			ID:          *i.ID,
			Name:        *i.DefaultLocale.Name,
			Version:     *i.Version,
			Enabled:     *i.Enabled,
			Description: *i.DefaultLocale.Description,
			// mozilla redirects to the addon page
			AddonURL:    fmt.Sprintf("https://addons.mozilla.org/en-US/firefox/addon/%s/", *i.ID),
			InstallDate: common.UnixMillis(*i.DateInstalled),

			DisabledByUser: *i.DisabledByUser,
			Author:         *i.DefaultLocale.Creator,
		}
	}

	return nil, nil
}
