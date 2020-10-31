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

package common

import (
	"encoding/json"
	"github.com/sandorex/ebd/errors"
	"time"
)

// ReadVersionFromJSON reads version field from JSON
func ReadVersionFromJSON(key string, data map[string]*json.RawMessage) (int, error) {
	versionFieldRaw, ok := data[key]
	if !ok {
		return 0, errors.ErrJSONMissingKeys([]string{key})
	}

	var versionInt int
	err := json.Unmarshal(*versionFieldRaw, &versionInt)
	if err != nil {
		return 0, err
	}

	return versionInt, nil
}

// UnixMillis converts milliseconds from epoch to Time
func UnixMillis(epoch int64) time.Time {
	return time.Unix(0, epoch/1000000)
}
