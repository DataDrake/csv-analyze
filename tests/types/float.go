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
	"fmt"
	"github.com/DataDrake/csv-analyze/tests"
	"io"
	"strconv"
)

const floatResultFormat = "\tFloat: %s\n"

// FloatTest checks if a column contains only floating point numbers and identifies their size
type FloatTest struct {
	failed64 bool
	failed32 bool
}

// NewFloatTest returns a fresh FloatTest
func NewFloatTest() tests.Test {
	return &FloatTest{false, false}
}

// Run attempts to convert the current cell to a float64, then ckecks if it can be a float32
func (f *FloatTest) Run(cell string) {
	if f.failed64 {
		return
	}
	val, err := strconv.ParseFloat(cell, 64)
	if err != nil {
		f.failed64 = true
		return
	}
	if f.failed32 {
		return
	}
	f.failed32 = val != float64(float32(val))
}

// Passed returns true if and only if all entries were floats
func (f *FloatTest) Passed() bool {
	return !f.failed64
}

// PrintResult indicates either failure or the smallest applicable float type
func (f *FloatTest) PrintResult(out io.Writer) {
	if f.failed64 {
		fmt.Fprintf(out, floatResultFormat, "fail")
		return
	}
	if !f.failed32 {
		fmt.Fprintf(out, floatResultFormat, "float32")
		return
	}
	fmt.Fprintf(out, floatResultFormat, "float64")
	return
}
