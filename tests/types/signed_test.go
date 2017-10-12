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

// TestNewSignedTest verifies the starting state of an SignedTest
func TestNewSignedTest(t *testing.T) {
	st := NewSignedTest().(*SignedTest)
	if st.mag != 0 {
		t.Errorf("Magnitude should be '%d', found: '%d'", 0, st.mag)
	}
	if st.failed {
		t.Error("Fail should not be true")
	}
}

// TestSignedTestRun verifies the correct operation of the run command
func TestNewSignedTestRun(t *testing.T) {
	st := NewSignedTest().(*SignedTest)
	st.Run("-512")
	if !st.Passed() {
		t.Error("Should have passed")
	}
	if st.mag != 512 {
		t.Errorf("Magnitude should be '%d', found: '%d'", 512, st.mag)
	}
	st.Run("1234")
	if !st.Passed() {
		t.Error("Should have passed")
	}
	if st.mag != 1234 {
		t.Errorf("Magnitude should be '%d', found: '%d'", 1234, st.mag)
	}
	st.Run("534")
	if !st.Passed() {
		t.Error("Should have passed")
	}
	if st.mag != 1234 {
		t.Errorf("Magnitude should be '%d', found: '%d'", 1234, st.mag)
	}
	st.Run("2345")
	if !st.Passed() {
		t.Error("Should have passed")
	}
	if st.mag != 2345 {
		t.Errorf("Magnitude should be '%d', found: '%d'", 2345, st.mag)
	}
	st.Run("bob")
	if st.Passed() {
		t.Error("Should have failed")
	}
	st.Run("1234")
	if st.Passed() {
		t.Error("Should have failed")
	}
}

// TestSignedTestPrintResult verifies the correct operation of the run command
func TestNewSignedTestPrintResult(t *testing.T) {
	buff := bytes.NewBuffer(make([]byte, 10))
	buff.Reset()
	ut := NewSignedTest()
	ut.Run("64")
	ut.PrintResult(buff)
	if buff.String() != fmt.Sprintf(signedResultFormat, "int8") {
		t.Errorf("Should have been int8, found: '%s'", buff.String())
	}
	buff.Reset()
	ut.Run("12800")
	ut.PrintResult(buff)
	if buff.String() != fmt.Sprintf(signedResultFormat, "int16") {
		t.Errorf("Should have been int16, found: '%s'", buff.String())
	}
	buff.Reset()
	ut.Run("128000")
	ut.PrintResult(buff)
	if buff.String() != fmt.Sprintf(signedResultFormat, "int32") {
		t.Errorf("Should have been int32, found: '%s'", buff.String())
	}
	buff.Reset()
	ut.Run("128000000000")
	ut.PrintResult(buff)
	if buff.String() != fmt.Sprintf(signedResultFormat, "int64") {
		t.Errorf("Should have been int64, found: '%s'", buff.String())
	}
	buff.Reset()
	ut.Run("bob")
	ut.PrintResult(buff)
	if buff.String() != fmt.Sprintf(signedResultFormat, "fail") {
		t.Errorf("Should have failed, found: '%s'", buff.String())
	}
}
