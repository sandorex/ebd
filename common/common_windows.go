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

// +build windows

package common

import (
	"github.com/sandorex/ebd/profile"
	"golang.org/x/sys/windows"
	"os"
)

// isFileOpen checks if file is open in another process
func isFileOpen(path string) (bool, error) {
	err := os.Rename(path, path)

	// only check for this specific error (code 32)
	if errno := err.(*os.LinkError).Unwrap(); errno == windows.ERROR_SHARING_VIOLATION {
		return true, nil
	}

	// any other error is propagated
	return false, err
}

func ReadProfileStateFromLockfile(path string) (profile.State, error) {
	result, err := isFileOpen(path)

	// if it does not exist it must mean it's running
	if os.IsNotExist(err) {
		return profile.StateClosed, nil
	}

	// file is open so it must be running
	if result {
		return profile.StateRunning, nil
	}

	// any other case is unknown state
	return profile.StateUnknown, err
}
