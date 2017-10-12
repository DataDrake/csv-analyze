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
	"strings"
)

const boolResultFormat = "\tBoolean: True - %v, False - %v \n"

var possibleTrue = []string{"true", "t", "yes", "y"}
var possibleFalse = []string{"false", "f", "no", "n"}

func contains(vals []string, match string) bool {
	for _, val := range vals {
		if val == match {
			return true
		}
	}
	return false
}

// BooleanTest checks if a column contains only boolstamps and identifies their layout
type BooleanTest struct {
	trues  []string
	falses []string
	failed bool
}

// NewBooleanTest returns a fresh BooleanTest
func NewBooleanTest() tests.Test {
	return &BooleanTest{make([]string, 0), make([]string, 0), false}
}

// Run attempts to convert the current cell to many different bool types
func (t *BooleanTest) Run(cell string) {
	if !t.Passed() {
		return
	}
	cell = strings.ToLower(strings.TrimSpace(cell))
	for _, match := range possibleTrue {
		if cell == match {
			if !contains(t.trues, match) {
				t.trues = append(t.trues, match)
			}
			return
		}
	}
	for _, match := range possibleFalse {
		if cell == match {
			if !contains(t.falses, match) {
				t.falses = append(t.falses, match)
			}
			return
		}
	}
	if v, err := strconv.ParseUint(cell, 10, 64); err == nil {
		if v == 0 {
			if !contains(t.falses, cell) {
				t.falses = append(t.falses, cell)
			}
			return
		}
		if v == 1 {
			if !contains(t.trues, cell) {
				t.trues = append(t.trues, cell)
			}
			return
		}
	}
	if v, err := strconv.ParseFloat(cell, 64); err == nil {
		if v == 0 {
			if !contains(t.falses, cell) {
				t.falses = append(t.falses, cell)
			}
			return
		}
		if v == 1 {
			if !contains(t.trues, cell) {
				t.trues = append(t.trues, cell)
			}
			return
		}
	}
	t.failed = true
	return
}

// Passed returns true if and only if all entries were bools
func (t *BooleanTest) Passed() bool {
	return !t.failed
}

// PrintResult indicates either failure or the detected bool types
func (t *BooleanTest) PrintResult(out io.Writer) {
	if !t.Passed() {
		fmt.Fprintf(out, boolResultFormat, "fail", "fail")
		return
	}
	fmt.Fprintf(out, boolResultFormat, t.trues, t.falses)
	return
}
