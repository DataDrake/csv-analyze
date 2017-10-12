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

// TestNewUnsignedTest verifies the starting state of an UnsignedTest
func TestNewUnsignedTest(t *testing.T) {
	ut := NewUnsignedTest().(*UnsignedTest)
	if ut.max != 0 {
		t.Errorf("Max should be '%d', found: '%d'", 0, ut.max)
	}
	if ut.failed {
		t.Error("Fail should not be true")
	}
}

// TestUnsignedTestRun verifies the correct operation of the run command
func TestNewUnsignedTestRun(t *testing.T) {
	ut := NewUnsignedTest().(*UnsignedTest)
	ut.Run("1234")
	if !ut.Passed() {
		t.Error("Should have passed")
	}
	if ut.max != 1234 {
		t.Errorf("Max should be '%d', found: '%d'", 1234, ut.max)
	}
	ut.Run("534")
	if !ut.Passed() {
		t.Error("Should have passed")
	}
	if ut.max != 1234 {
		t.Errorf("Max should be '%d', found: '%d'", 1234, ut.max)
	}
	ut.Run("2345")
	if !ut.Passed() {
		t.Error("Should have passed")
	}
	if ut.max != 2345 {
		t.Errorf("Max should be '%d', found: '%d'", 2345, ut.max)
	}
	ut.Run("bob")
	if ut.Passed() {
		t.Error("Should have failed")
	}
	ut.Run("1234")
	if ut.Passed() {
		t.Error("Should have failed")
	}
}

// TestUnsignedTestPrintResult verifies the correct operation of the run command
func TestNewUnsignedTestPrintResult(t *testing.T) {
	buff := bytes.NewBuffer(make([]byte, 10))
	buff.Reset()
	ut := NewUnsignedTest()
	ut.Run("128")
	ut.PrintResult(buff)
	if buff.String() != fmt.Sprintf(unsignedResultFormat, "uint8") {
		t.Errorf("Should have been uint8, found: '%s'", buff.String())
	}
	buff.Reset()
	ut.Run("12800")
	ut.PrintResult(buff)
	if buff.String() != fmt.Sprintf(unsignedResultFormat, "uint16") {
		t.Errorf("Should have been uint16, found: '%s'", buff.String())
	}
	buff.Reset()
	ut.Run("128000")
	ut.PrintResult(buff)
	if buff.String() != fmt.Sprintf(unsignedResultFormat, "uint32") {
		t.Errorf("Should have been uint32, found: '%s'", buff.String())
	}
	buff.Reset()
	ut.Run("128000000000")
	ut.PrintResult(buff)
	if buff.String() != fmt.Sprintf(unsignedResultFormat, "uint64") {
		t.Errorf("Should have been uint64, found: '%s'", buff.String())
	}
	buff.Reset()
	ut.Run("bob")
	ut.PrintResult(buff)
	if buff.String() != fmt.Sprintf(unsignedResultFormat, "fail") {
		t.Errorf("Should have failed, found: '%s'", buff.String())
	}
}
