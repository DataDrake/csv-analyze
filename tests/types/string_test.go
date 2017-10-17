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

// TestNewStringTest verifies the starting state of an StringTest
func TestNewStringTest(t *testing.T) {
	st := NewStringTest().(*StringTest)
	if st.Passed() {
		t.Error("Should not have passed")
	}
	if st.max != 0 {
		t.Error("'max' should be 0")
	}
}

// TestStringTestRun verifies the correct operation of the run command
func TestNewStringTestRun(t *testing.T) {
	st := NewStringTest().(*StringTest)
	st.Run("hello")
	if !st.Passed() {
		t.Error("Should have passed")
	}
	st.Run("hello1234")
	if !st.Passed() {
		t.Error("Should have passed")
	}
	st.Run("hello1234\n")
	if !st.Passed() {
		t.Error("Should have passed")
	}
	st.Run("☃💩")
	if !st.Passed() {
		t.Error("Should have passed")
	}
}

// TestStringTestPrintResult verifies the correct operation of the run command
func TestNewStringTestPrintResult(t *testing.T) {
	buff := bytes.NewBuffer(make([]byte, 10))
	buff.Reset()
	st := NewStringTest()
	st.PrintResult(buff)
	if buff.String() != fmt.Sprintf(stringResultFormat, "Fail", 0) {
		t.Errorf("Should have been 'Fail' and '0', found: '%s'", buff.String())
	}
	buff.Reset()
	st.Run("hello")
	st.PrintResult(buff)
	if buff.String() != fmt.Sprintf(stringResultFormat, "Alpha", 5) {
		t.Errorf("Should have been 'Alpha' and '5', found: '%s'", buff.String())
	}
	buff.Reset()
	st.Run("hello1234")
	st.PrintResult(buff)
	if buff.String() != fmt.Sprintf(stringResultFormat, "Alphanumeric", 9) {
		t.Errorf("Should have been 'Alphanumeric' and '9', found: '%s'", buff.String())
	}
	buff.Reset()
	st.Run("hello1234\n")
	st.PrintResult(buff)
	if buff.String() != fmt.Sprintf(stringResultFormat, "ASCII", 10) {
		t.Errorf("Should have been 'ASCII' and '10', found: '%s'", buff.String())
	}
	buff.Reset()
	st.Run("☃💩")
	st.PrintResult(buff)
	if buff.String() != fmt.Sprintf(stringResultFormat, "UTF-8", 10) {
		t.Errorf("Should have been 'UTF-8' and '10', found: '%s'", buff.String())
	}
}
