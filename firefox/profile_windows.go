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

package firefox

import (
	"github.com/sandorex/ebd/profile"
	"golang.org/x/sys/windows"
	"os"
	"path"
)

// GetProfileState reads profile state by checking if the lockfile is open in
// another process, if it is then the profile is running, if it isn't then it's
// closed
//
// NOTE: THE LOCKFILE IS NOT DELETED WHEN FIREFOX CLOSES
func (p Profile) GetProfileState() (profile.State, error) {
	lockfile := path.Join(p.path, FileLockfile)

	// rename inplace
	err := os.Rename(lockfile, lockfile)

	// if there was no error it's probably not running
	if err == nil {
		return profile.StateClosed, nil
	}

	// it's open in another process so it must be running
	if errno := err.(*os.LinkError).Unwrap(); errno == windows.ERROR_SHARING_VIOLATION {
		return profile.StateRunning, nil
	}

	// any other case is unknown state
	return profile.StateUnknown, err
}
