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
	"testing"
)

func TestFFPIDExtraction(t *testing.T) {
	tables := []struct {
		linkTarget string
		pid        int
	}{
		{"0.0.0.0:+1234", 1234},
		{"127.0.0.1:+9999", 9999},
	}

	for _, table := range tables {
		pid, err := extractPID(table.linkTarget)
		if pid != table.pid || err != nil {
			t.Errorf("pid extracted from %q resulted in %d instead of %d\nerror: %v", table.linkTarget, pid, table.pid, err)
		}
	}
}
