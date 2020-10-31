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

package main

import (
	"fmt"
	"github.com/sandorex/ebd/firefox"
)

func main() {
	p := firefox.NewProfile("/home/sandorex/.mozilla/firefox/uxwfkoxd.resync-new", "", false)
	// p := chromium.NewProfile("/home/sandorex/.config/chromium/Default")

	if p.IsProfileValid() {
		fmt.Printf("The profile %q at %q is valid\n", p.GetProfileName(), p.GetProfilePath())
	} else {
		fmt.Printf("The profile %q at %q is not valid\n", p.GetProfileName(), p.GetProfilePath())
	}

	state, err := p.GetProfileState()
	fmt.Printf("The profile state %v, %v\n", state, err)
	if p.IsProfileRunning() {
		fmt.Println("The profile is running")
	} else {
		fmt.Println("The profile is not running")
	}
}
