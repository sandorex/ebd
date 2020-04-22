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

// Package profile contains base profile api
package profile

// State represents state of a profile
type State int8

const (
	// StateClosed means the profile is not in use
	StateClosed State = iota

	// StateRunning means the profile is in use
	StateRunning State = iota

	// StateUnknown means the state of profile is either crashed on corrupted and
	// as such writing to it is not recommended, reading also isn't recommended
	// but will not damage the data, the read that just may be wrong
	StateUnknown State = iota
)

// Profile is interface to a profile of any browser
type Profile interface {
	GetProfileName() string
	GetProfilePath() string
	GetProfileState() State
	IsProfileValid() bool
	IsProfileRunning() bool
	OpenReader()
	OpenWriter()
}
