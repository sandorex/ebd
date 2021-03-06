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

package firefox

const (
	// FileProfiles contains profile paths and names
	FileProfiles = "profiles.ini"

	// FileSessionStore contains data from last session
	FileSessionStore = "sessionstore.jsonlz4"

	// FileExtensions contains extensions and addons installed
	FileExtensions = "extensions.json"

	// FilePlacesDatabase contains history and bookmarks
	FilePlacesDatabase = "places.sqlite"

	// FileCookiesDatabase contains cookies
	FileCookiesDatabase = "cookies.sqlite"

	// FileSignedInUser contains username and email of the signed in user
	FileSignedInUser = "signedInUser.json"

	// FileContainers contains containers (exists even if the extension is not
	// installed)
	FileContainers = "containers.json"
)
