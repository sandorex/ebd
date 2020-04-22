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

// +build !windows

package chromium

import (
	"github.com/sandorex/ebd/common"
	"github.com/sandorex/ebd/profile"
	"path/filepath"
	"strconv"
	"strings"
)

// extractPID extracts PID from string using format that Chromium lockfile uses
// 'HOSTNAME-PID'
func extractPID(linkTarget string) (int, error) {
	index := strings.LastIndex(linkTarget, "-")
	if index == -1 {
		// '-' has not been found
		return -1, nil
	}

	return strconv.Atoi(linkTarget[index+1:])
}

// GetProfileState reads profile state by reading the lockfile link if the
// lockfile doesn't exist the profile is closed, if it does but the PID is not
// a valid process then the profile has crashed
func (p Profile) GetProfileState() (profile.State, error) {
	return common.ReadProfileStateFromLockfile(filepath.Join(filepath.Dir(p.path), FileLockfile), extractPID)
}
