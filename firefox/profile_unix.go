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

package firefox

import (
	"github.com/sandorex/ebd/profile"
	"os"
	"path"
	"strconv"
	"strings"
	"syscall"
)

// extractPID extracts PID from string using format that Firefox lockfile uses
// 'ADDR:+PID' where ADDR is an IP4 address
func extractPID(linkTarget string) (int, error) {
	index := strings.LastIndex(linkTarget, ":+")
	if index == -1 {
		// ':+' has not been found
		return -1, nil
	}

	return strconv.Atoi(linkTarget[index+2:])
}

// GetProfileState reads profile state by reading the lockfile link if the
// lockfile doesn't exist the profile is closed, if it does but the PID is not
// a valid process then the profile has crashed
func (p Profile) GetProfileState() (profile.State, error) {
	linkTarget, err := os.Readlink(path.Join(p.path, FileLockfile))

	// if it does not exist it probably isn't running
	if os.IsNotExist(err) {
		return profile.StateClosed, nil
	}

	// any other error is unknown state
	if err != nil {
		return profile.StateUnknown, err
	}

	pid, err := extractPID(linkTarget)
	if pid == -1 || err != nil {
		// cannot extract pid from the target, something is definitely wrong
		return profile.StateUnknown, err
	}

	// WIP
	process, err := os.FindProcess(pid)
	if err != nil {
		return profile.StateUnknown, err
	}

	err = process.Signal(syscall.Signal(0))
	return profile.StateUnknown, err
}
