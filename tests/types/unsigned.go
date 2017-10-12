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

const unsignedResultFormat = "\tUnsigned: %s\n"

// UnsignedTest checks if a column contains only unsigned integers and identifies their size
type UnsignedTest struct {
	max    uint64
	failed bool
}

// NewUnsignedTest returns a fresh UnsignedTest
func NewUnsignedTest() tests.Test {
	return &UnsignedTest{0, false}
}

// Run attempts to convert the current cell to an integer, storing the maximum if found
func (u *UnsignedTest) Run(cell string) {
	if u.failed {
		return
	}
	val, err := strconv.ParseUint(cell, 10, 64)
	if err != nil {
		u.failed = true
		return
	}
	if val > u.max {
		u.max = val
	}
}

// Passed returns true if and only if all entries were unsigned integers
func (u *UnsignedTest) Passed() bool {
	return !u.failed
}

// PrintResult indicates either failure or the smallest applicable unsigned type
func (u *UnsignedTest) PrintResult(out io.Writer) {
	if u.failed {
		fmt.Fprintf(out, unsignedResultFormat, "fail")
		return
	}
	switch {
	case u.max <= math.MaxUint8:
		fmt.Fprintf(out, unsignedResultFormat, "uint8")
	case u.max <= math.MaxUint16:
		fmt.Fprintf(out, unsignedResultFormat, "uint16")
	case u.max <= math.MaxUint32:
		fmt.Fprintf(out, unsignedResultFormat, "uint32")
	default:
		fmt.Fprintf(out, unsignedResultFormat, "uint64")
	}
	return
}
