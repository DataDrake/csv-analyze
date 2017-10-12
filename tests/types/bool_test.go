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

// TestNewBooleanTest verifies the starting state of an BooleanTest
func TestNewBooleanTest(t *testing.T) {
	bt := NewBooleanTest().(*BooleanTest)
	if !bt.Passed() {
		t.Error("Should have passed")
	}
	if len(bt.trues) != 0 {
		t.Error("trues should be empty")
	}
	if len(bt.falses) != 0 {
		t.Error("falses should be empty")
	}
}

// TestBooleanTestRun verifies the correct operation of the run command
func TestNewBooleanTestRun(t *testing.T) {
	bt := NewBooleanTest().(*BooleanTest)
	bt.Run("true")
	if !bt.Passed() {
		t.Error("Should have passed")
	}
	if len(bt.trues) != 1 {
		t.Error("'trues' should have exaclty 1 entry")
	}
	if !contains(bt.trues, "true") {
		t.Error("'true' should have been found")
	}
	bt.Run("False")
	if !bt.Passed() {
		t.Error("Should have passed")
	}
	if len(bt.trues) != 1 {
		t.Error("'trues' should have exaclty 1 entry")
	}
	if !contains(bt.trues, "true") {
		t.Error("'true' should have been found")
	}
	if len(bt.falses) != 1 {
		t.Error("'falses' should have exaclty 1 entry")
	}
	if !contains(bt.falses, "false") {
		t.Error("'true' should have been found")
	}
	bt.Run("0")
	if !bt.Passed() {
		t.Error("Should have passed")
	}
	if len(bt.trues) != 1 {
		t.Error("'trues' should have exaclty 1 entry")
	}
	if !contains(bt.trues, "true") {
		t.Error("'true' should have been found")
	}
	if len(bt.falses) != 2 {
		t.Error("'falses' should have exaclty 1 entry")
	}
	if !contains(bt.falses, "false") {
		t.Error("'true' should have been found")
	}
	if !contains(bt.falses, "0") {
		t.Error("'true' should have been found")
	}
	bt.Run("1")
	if !bt.Passed() {
		t.Error("Should have passed")
	}
	if len(bt.trues) != 2 {
		t.Error("'trues' should have exaclty 1 entry")
	}
	if !contains(bt.trues, "true") {
		t.Error("'true' should have been found")
	}
	if !contains(bt.trues, "1") {
		t.Error("'true' should have been found")
	}
	if len(bt.falses) != 2 {
		t.Error("'falses' should have exaclty 1 entry")
	}
	if !contains(bt.falses, "false") {
		t.Error("'true' should have been found")
	}
	if !contains(bt.falses, "0") {
		t.Error("'true' should have been found")
	}
	bt.Run("1.0")
	if !bt.Passed() {
		t.Error("Should have passed")
	}
	if len(bt.trues) != 3 {
		t.Error("'trues' should have exaclty 1 entry")
	}
	if !contains(bt.trues, "true") {
		t.Error("'true' should have been found")
	}
	if !contains(bt.trues, "1") {
		t.Error("'true' should have been found")
	}
	if !contains(bt.trues, "1.0") {
		t.Error("'true' should have been found")
	}
	if len(bt.falses) != 2 {
		t.Error("'falses' should have exaclty 1 entry")
	}
	if !contains(bt.falses, "false") {
		t.Error("'true' should have been found")
	}
	if !contains(bt.falses, "0") {
		t.Error("'true' should have been found")
	}
	bt.Run("0.0")
	if !bt.Passed() {
		t.Error("Should have passed")
	}
	if len(bt.trues) != 3 {
		t.Error("'trues' should have exaclty 1 entry")
	}
	if !contains(bt.trues, "true") {
		t.Error("'true' should have been found")
	}
	if !contains(bt.trues, "1") {
		t.Error("'true' should have been found")
	}
	if !contains(bt.trues, "1.0") {
		t.Error("'true' should have been found")
	}
	if len(bt.falses) != 3 {
		t.Error("'falses' should have exaclty 1 entry")
	}
	if !contains(bt.falses, "false") {
		t.Error("'true' should have been found")
	}
	if !contains(bt.falses, "0") {
		t.Error("'true' should have been found")
	}
	if !contains(bt.falses, "0.0") {
		t.Error("'true' should have been found")
	}
	bt.Run("bob")
	if bt.Passed() {
		t.Error("Should not have passed")
	}
	bt.Run("dave")
	if bt.Passed() {
		t.Error("Should not have passed")
	}
}

// TestBooleanTestPrintResult verifies the correct operation of the run command
func TestNewBooleanTestPrintResult(t *testing.T) {
	buff := bytes.NewBuffer(make([]byte, 10))
	buff.Reset()
	bt := NewBooleanTest()
	bt.Run("true")
	bt.PrintResult(buff)
	if buff.String() != fmt.Sprintf(boolResultFormat, []string{"true"}, []string{}) {
		t.Errorf("Should have been (true) (), found: '%s'", buff.String())
	}
	buff.Reset()
	bt.Run("False")
	bt.PrintResult(buff)
	if buff.String() != fmt.Sprintf(boolResultFormat, []string{"true"}, []string{"false"}) {
		t.Errorf("Should have been (true) (false), found: '%s'", buff.String())
	}
	buff.Reset()
	bt.Run("dave")
	bt.PrintResult(buff)
	if buff.String() != fmt.Sprintf(boolResultFormat, "fail", "fail") {
		t.Errorf("Should have failed, found: '%s'", buff.String())
	}
}
