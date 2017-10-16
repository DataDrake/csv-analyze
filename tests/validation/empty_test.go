//
// Copyright 2017 Bryan T. Meyers <bmeyers@datadrake.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package validation

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"testing"
)

// TestNewEmptySuite verifies the starting state of an EmptySuite
func TestNewEmptySuite(t *testing.T) {
	et := NewEmptySuite().(*EmptySuite)
	if et == nil {
		t.Error("EmptySuite should not be nil")
	}
}

const testCSV = "1,2,3,4\n5,,6,7\n1,2,\n"
const testCSV2 = "1,2,3,4\n5,,6,7\n"

// TestNewEmptySuiteRun verifies the correct operation of the run command
func TestNewEmptySuiteRun(t *testing.T) {
	dst := bytes.NewBuffer(make([]byte, 10))
	dst.Reset()
	src := csv.NewReader(bytes.NewBuffer([]byte(testCSV)))
	et := NewEmptySuite()
	err := et.Run(src, dst)
	if dst.String() != fmt.Sprintf(emptyResultFormat, 1, 1) {
		t.Errorf("Should have been (1) (1), found: '%s'", dst.String())
	}
	if err == nil {
		t.Error("There should have been an error")
	}
}

// TestNewEmptySuiteRun2 verifies the correct operation of the run command
func TestNewEmptySuiteRun2(t *testing.T) {
	dst := bytes.NewBuffer(make([]byte, 10))
	dst.Reset()
	src := csv.NewReader(bytes.NewBuffer([]byte(testCSV2)))
	et := NewEmptySuite()
	err := et.Run(src, dst)
	if dst.String() != fmt.Sprintf(emptyResultFormat, 1, 1) {
		t.Errorf("Should have been (1) (1), found: '%s'", dst.String())
	}
	if err != nil {
		t.Error("There should not have been an error")
	}
}
