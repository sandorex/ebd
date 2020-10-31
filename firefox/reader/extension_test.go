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
	"fmt"
	"github.com/google/go-cmp/cmp"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

var extensionDatapath = filepath.Join("datafiles", "extensions.json")

func TestFFExtensionsDefault(t *testing.T) {
	data := map[int]Container{}

	for version := range extensionVersions {
		path := filepath.Join(extensionDatapath, fmt.Sprintf("v%dd.json", version))

		file, err := os.Open(path)
		if err != nil {
			t.Errorf("could not open the file %q\n%s\n", path, err)
			continue
		}

		defer file.Close()

		bytes, err := ioutil.ReadAll(file)
		if err != nil {
			t.Errorf("could not read the file %q\n%v\n", path, err)
			continue
		}

		parsedData, err := ParseExtensions(bytes)
		if err != nil {
			t.Errorf("error parsing %q\n%v\n", path, err)
			continue
		}

		if !cmp.Equal(data, parsedData) {
			t.Errorf("parsed result in %q does not match predefined data\n%s\n", path, cmp.Diff(data, parsedData))
		}

		t.Logf("output: %v", parsedData)
	}
}

// func TestFFExtensionsModified(t *testing.T) {
// 	data := map[int]Container{
// 		1: Container{
// 			ID:    1,
// 			Name:  "1",
// 			Icon:  "fingerprint",
// 			Color: "purple",
// 		},
// 		3: Container{
// 			ID:    3,
// 			Name:  "3",
// 			Icon:  "circle",
// 			Color: "green",
// 		},
// 		6: Container{
// 			ID:    6,
// 			Name:  "6",
// 			Icon:  "fence",
// 			Color: "toolbar",
// 		},
// 	}

// 	for version := range containerVersions {
// 		path := filepath.Join(extensionDatapath, fmt.Sprintf("v%dm.json", version))

// 		file, err := os.Open(path)
// 		if err != nil {
// 			t.Errorf("could not open the file %q\n%v\n", path, err)
// 			continue
// 		}

// 		defer file.Close()

// 		bytes, err := ioutil.ReadAll(file)
// 		if err != nil {
// 			t.Errorf("could not read the file %q\n%v\n", path, err)
// 			continue
// 		}

// 		parsedData, err := ParseContainers(bytes)
// 		if err != nil {
// 			t.Errorf("error parsing %q\n%v\n", path, err)
// 			continue
// 		}

// 		if !cmp.Equal(data, parsedData) {
// 			t.Errorf("parsed result in %q does not match predefined data\n%s\n", path, cmp.Diff(data, parsedData))
// 		}
// 	}
// }
