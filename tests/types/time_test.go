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

package types

import (
	"bytes"
	"fmt"
	"testing"
)

// TestNewTimeTest verifies the starting state of an TimeTest
func TestNewTimeTest(t *testing.T) {
	tt := NewTimeTest().(*TimeTest)
	if !tt.Passed() {
		t.Error("Should have passed")
	}
}

// TestTimeTestRun verifies the correct operation of the run command
func TestNewTimeTestRun(t *testing.T) {
	tt := NewTimeTest().(*TimeTest)
	tt.Run("7:01AM")
	if !tt.Passed() {
		t.Error("Should have passed")
	}
	if !tt.cases[10].found {
		t.Error("Kitchen should have been found")
	}
	tt.Run("2017-10-05T23:46:57-03:00")
	if !tt.Passed() {
		t.Error("Should have passed")
	}
	if !tt.cases[8].found {
		t.Error("RFC3339 should have been found")
	}
	if !tt.cases[10].found {
		t.Error("Kitchen should have been found")
	}
	tt.Run("bob")
	if tt.Passed() {
		t.Error("Should not have passed")
	}
	tt.Run("dave")
	if tt.Passed() {
		t.Error("Should not have passed")
	}
}

// TestTimeTestPrintResult verifies the correct operation of the run command
func TestNewTimeTestPrintResult(t *testing.T) {
	buff := bytes.NewBuffer(make([]byte, 10))
	buff.Reset()
	tt := NewTimeTest()
	tt.Run("7:01AM")
	tt.PrintResult(buff)
	if buff.String() != fmt.Sprintf(timeResultFormat, []string{"Kitchen"}) {
		t.Errorf("Should have been (Kitcken), found: '%s'", buff.String())
	}
	buff.Reset()
	tt.Run("2017-10-05T23:46:57-03:00")
	tt.PrintResult(buff)
	if buff.String() != fmt.Sprintf(timeResultFormat, []string{"RFC3339", "RFC3339Nano","Kitchen"}) {
		t.Errorf("Should have been (RFC3339,RFC3339Nano,Kitchen), found: '%s'", buff.String())
	}
	buff.Reset()
	tt.Run("dave")
	tt.PrintResult(buff)
	if buff.String() != fmt.Sprintf(timeResultFormat, "fail") {
		t.Errorf("Should have failed, found: '%s'", buff.String())
	}
}
