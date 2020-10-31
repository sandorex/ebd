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

package errors

import (
	"errors"
	"fmt"
)

// ErrUnsupportedSchemaVersion represents error when unsupported schema version
// is read
type ErrUnsupportedSchemaVersion struct {
	// Version is the version found
	Version int

	// LastSupportedVersion is the last version that is compatible
	LastSupportedVersion *int
}

func (e ErrUnsupportedSchemaVersion) Error() string {
	if e.LastSupportedVersion != nil {
		return fmt.Sprintf("Unsupported schema version %d (last supported version %d)", e.Version, e.LastSupportedVersion)
	}

	return fmt.Sprintf("Unsupported schema version %d", e.Version)
}

// ErrJSONMissingKeys represents error when keys are missing when reading JSON
// at least one of these should be defined
type ErrJSONMissingKeys []string

func (s ErrJSONMissingKeys) Error() string {
	return fmt.Sprintf("None of these keys were found in JSON: %v", []string(s))
}

// ErrCorruptedData means the data is wrong in some way that is safe to assume
// it was changed manually or corrupted
var ErrCorruptedData = errors.New("The data is probably corrupted")
