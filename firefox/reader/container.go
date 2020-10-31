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
	"github.com/asaskevich/govalidator"
	"github.com/sandorex/ebd/common"
	"github.com/sandorex/ebd/errors"
)

// Container represents Firefox session container
type Container struct {
	ID    int
	Name  string
	Icon  string
	Color string
}

// ParseContainers parses containers from the file
func ParseContainers(data []byte) (map[int]Container, error) {
	var v map[string]*json.RawMessage
	err := json.Unmarshal(data, &v)
	if err != nil {
		return nil, err
	}

	version, err := common.ReadVersionFromJSON("version", v)
	if err != nil {
		return nil, err
	}

	fn, ok := containerVersions[version]
	if !ok {
		return nil, errors.ErrUnsupportedSchemaVersion{Version: version}
	}

	return fn(v)
}

var containerVersions = map[int](func(data map[string]*json.RawMessage) (map[int]Container, error)){
	4: parseContainersV4,
}

func parseContainersV4(data map[string]*json.RawMessage) (map[int]Container, error) {
	rawIdentities, ok := data["identities"]
	if !ok {
		return nil, errors.ErrJSONMissingKeys([]string{"identities"})
	}

	var v []struct {
		ID           *int    `json:"userContextId" valid:"required"`
		Name         *string `json:"name" valid:"optional"`
		Icon         *string `json:"icon" valid:"required"`
		Color        *string `json:"color" valid:"required"`
		Public       *bool   `json:"public" valid:"required"`
		Localization *string `json:"l10nID" valid:"optional"`
	}
	err := json.Unmarshal(*rawIdentities, &v)
	if err != nil {
		return nil, err
	}

	// convert to common type
	containers := map[int]Container{}
	for _, i := range v {
		_, err := govalidator.ValidateStruct(i)
		if err != nil {
			return nil, err
		}

		if *i.Public {
			container := Container{}
			container.ID = *i.ID
			container.Icon = *i.Icon
			container.Color = *i.Color

			// using localization string as the name if the name is not set
			if i.Name != nil {
				container.Name = *i.Name
			} else if i.Localization != nil {
				container.Name = *i.Localization
			} else {
				return nil, errors.ErrJSONMissingKeys([]string{"name", "l10nID"})
			}

			containers[container.ID] = container
		}
	}

	return containers, nil
}
