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

// TestNewFloatTest verifies the starting state of an FloatTest
func TestNewFloatTest(t *testing.T) {
	ft := NewFloatTest().(*FloatTest)
	if ft.failed64 {
		t.Error("failed64 should not be true")
	}
	if ft.failed32 {
		t.Error("failed32 should not be true")
	}
}

// TestFloatTestRun verifies the correct operation of the run command
func TestNewFloatTestRun(t *testing.T) {
	ft := NewFloatTest().(*FloatTest)
	ft.Run("1234.0")
	if !ft.Passed() {
		t.Error("Should have passed")
	}
	if ft.failed32 {
		t.Error("failed32 should not be true")
	}
	ft.Run("1E100")
	if !ft.Passed() {
		t.Error("Should have passed")
	}
	if !ft.failed32 {
		t.Error("failed32 should be true")
	}
	ft.Run("1E101")
	if !ft.Passed() {
		t.Error("Should have passed")
	}
	if !ft.failed32 {
		t.Error("failed32 should be true")
	}
	ft.Run("1E400")
	if ft.Passed() {
		t.Error("Should not have passed")
	}
	ft.Run("1E401")
	if ft.Passed() {
		t.Error("Should not have passed")
	}
}

// TestFloatTestPrintResult verifies the correct operation of the run command
func TestNewFloatTestPrintResult(t *testing.T) {
	buff := bytes.NewBuffer(make([]byte, 10))
	buff.Reset()
	ft := NewFloatTest()
	ft.Run("1234.0")
	ft.PrintResult(buff)
	if buff.String() != fmt.Sprintf(floatResultFormat, "float32") {
		t.Errorf("Should have been float32, found: '%s'", buff.String())
	}
	buff.Reset()
	ft.Run("1E100")
	ft.PrintResult(buff)
	if buff.String() != fmt.Sprintf(floatResultFormat, "float64") {
		t.Errorf("Should have been float64, found: '%s'", buff.String())
	}
	buff.Reset()
	ft.Run("1E400")
	ft.PrintResult(buff)
	if buff.String() != fmt.Sprintf(floatResultFormat, "fail") {
		t.Errorf("Should have failed, found: '%s'", buff.String())
	}
}
