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
	"time"
)

const timeResultFormat = "\tTime: %v\n"

type timeCase struct {
	name   string
	layout string
	found  bool
}

// TimeTest checks if a column contains only timestamps and identifies their layout
type TimeTest struct {
	cases  []timeCase
	failed bool
}

// NewTimeTest returns a fresh TimeTest
func NewTimeTest() tests.Test {
	return &TimeTest{
		[]timeCase{
			timeCase{"ANSIC", time.ANSIC, false},
			timeCase{"UnixDate", time.UnixDate, false},
			timeCase{"RubyDate", time.RubyDate, false},
			timeCase{"RFC822", time.RFC822, false},
			timeCase{"RFC822Z", time.RFC822Z, false},
			timeCase{"RFC850", time.RFC850, false},
			timeCase{"RFC1123", time.RFC1123, false},
			timeCase{"RFC1123Z", time.RFC1123Z, false},
			timeCase{"RFC3339", time.RFC3339, false},
			timeCase{"RFC3339Nano", time.RFC3339Nano, false},
			timeCase{"Kitchen", time.Kitchen, false},
			timeCase{"Stamp", time.Stamp, false},
			timeCase{"StampMilli", time.StampMilli, false},
			timeCase{"StampMicro", time.StampMicro, false},
			timeCase{"StampNano", time.StampNano, false},
			timeCase{"Apache", "02/Jan/2006:15:04:05 -0700", false},
		},
		false,
	}
}

// Run attempts to convert the current cell to many different time types
func (t *TimeTest) Run(cell string) {
	if !t.Passed() {
		return
	}
	t.failed = true
	for i, tc := range t.cases {
		_, err := time.Parse(tc.layout, cell)
		if err == nil {
			t.cases[i].found = true
			t.failed = false
		}
	}
	return
}

// Passed returns true if and only if all entries were times
func (t *TimeTest) Passed() bool {
	return !t.failed
}

// PrintResult indicates either failure or the detected time types
func (t *TimeTest) PrintResult(out io.Writer) {
	if !t.Passed() {
		fmt.Fprintf(out, timeResultFormat, "fail")
		return
	}
	matches := make([]string, 0)
	for _, tc := range t.cases {
		if tc.found {
			matches = append(matches, tc.name)
		}
	}
	fmt.Fprintf(out, timeResultFormat, matches)
	return
}
