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
	"math"
	"strconv"
)

const signedResultFormat = "\tSigned: %s\n"

// SignedTest checks if a column contains only signed integers and identifies their size
type SignedTest struct {
	mag    int64
	failed bool
}

// NewSignedTest returns a fresh SignedTest
func NewSignedTest() tests.Test {
	return &SignedTest{0, false}
}

// Run attempts to convert the current cell to an integer, storing the max magnitude if found
func (s *SignedTest) Run(cell string) {
	if s.failed {
		return
	}
	val, err := strconv.ParseInt(cell, 10, 64)
	if err != nil {
		s.failed = true
		return
	}
	if val < 0 {
		val *= -1
	}
	if val > s.mag {
		s.mag = val
	}
}

// Passed returns true if and only if all entries were signed integers
func (s *SignedTest) Passed() bool {
	return !s.failed
}

// PrintResult indicates either failure or the smallest applicable signed type
func (s *SignedTest) PrintResult(out io.Writer) {
	if s.failed {
		fmt.Fprintf(out, signedResultFormat, "fail")
		return
	}
	switch {
	case s.mag <= math.MaxInt8:
		fmt.Fprintf(out, signedResultFormat, "int8")
	case s.mag <= math.MaxInt16:
		fmt.Fprintf(out, signedResultFormat, "int16")
	case s.mag <= math.MaxInt32:
		fmt.Fprintf(out, signedResultFormat, "int32")
	default:
		fmt.Fprintf(out, signedResultFormat, "int64")
	}
	return
}
