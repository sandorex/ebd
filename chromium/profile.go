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

// Package chromium contains code specific to Chromium browser
package chromium

import (
	"github.com/sandorex/ebd/profile"
	"os"
	"path/filepath"
)

// Profile represents Chromium browser profile
//
// NOTE: it's not recommended to initialize the struct directly, use the
// NewProfile function
type Profile struct {
	name, path string
}

// NewProfile creates an instance of Profile but does not check if the profile
// points to a valid profile, run IsProfileValid to check
//
// NOTE: the name argument is optional
func NewProfile(path string) Profile {
	return Profile{
		path: path,
		name: filepath.Base(path),
	}
}

// GetProfileName returns name of the profile
func (p Profile) GetProfileName() string {
	return p.name
}

// GetProfilePath returns path to the profile (may not be valid)
func (p Profile) GetProfilePath() string {
	return p.path
}

// IsProfileValid check if profile instance points to a valid profile
//
// TODO: a more reliable way to check?
func (p Profile) IsProfileValid() bool {
	// check if each file that always exists in a profile exist
	for _, file := range []string{
		FilePreferences,
		FileHistoryDatabase,
		FileLoginData,
		FileWebData,
		FileCookiesDatabase,
		FileSecurePreferences,
		FileBookmarks,
	} {
		if _, err := os.Stat(filepath.Join(p.path, file)); os.IsNotExist(err) {
			return false
		}
	}

	return true
}

// IsProfileRunning checks if profile is running
//
// NOTE: unknown profile state is detected as if it was running to prevent data
// loss by writing to corrupted profiles
func (p Profile) IsProfileRunning() bool {
	state, err := p.GetProfileState()

	// detect as if it was running when the state is unknown
	if state != profile.StateClosed || err != nil {
		return true
	}

	return false
}

// OpenReader()
// OpenWriter()
