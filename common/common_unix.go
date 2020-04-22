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

package common

import (
	"github.com/sandorex/ebd/profile"
	"golang.org/x/sys/unix"
	"os"
	"syscall"
)

// isProcessRunning checks if process is still running
func isProcessRunning(pid int) bool {
	// if there is no errors it's still running
	return unix.Kill(pid, syscall.Signal(0)) == nil
}

// ReadProfileStateFromLockfile reads profile state from a lockfile link, param
// fn is the function that extracts PID from link target
//
// NOTE: this function is unix only
func ReadProfileStateFromLockfile(path string, fn func(string) (int, error)) (profile.State, error) {
	linkTarget, err := os.Readlink(path)

	// if it does not exist it probably isn't running
	if os.IsNotExist(err) {
		return profile.StateClosed, nil
	}

	// any other error is unknown state
	if err != nil {
		return profile.StateUnknown, err
	}

	// try to extract the pid
	pid, err := fn(linkTarget)
	if pid < 0 || err != nil {
		// cannot extract pid from the target, something is definitely wrong
		return profile.StateUnknown, err
	}

	// if the process is still running then it must still be open
	if isProcessRunning(pid) {
		return profile.StateRunning, nil
	}

	// all other cases are unknown state
	return profile.StateUnknown, err
}
