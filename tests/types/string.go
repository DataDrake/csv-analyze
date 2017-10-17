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
	"regexp"
)

const stringResultFormat = "\t%s, Max Length: %d\n"

var alpha *regexp.Regexp
var alphanumeric *regexp.Regexp
var ascii *regexp.Regexp

func init() {
	alpha = regexp.MustCompile("^[[:alpha:]]+$")
	alphanumeric = regexp.MustCompile("^[[:alnum:]]+$")
	ascii = regexp.MustCompile("^[[:ascii:]]+$")
}

type stringType struct {
	name   string
	regex  *regexp.Regexp
	failed bool
}

// StringTest checks if a column contains only boolstamps and identifies their layout
type StringTest struct {
	types []stringType
	max   uint64
}

// NewStringTest returns a fresh StringTest
func NewStringTest() tests.Test {
	return &StringTest{
		[]stringType{
			stringType{"Alpha", alpha, false},
			stringType{"Alphanumeric", alphanumeric, false},
			stringType{"ASCII", ascii, false},
		},
		0,
	}
}

// Run attempts to detect the character class of the current string
func (s *StringTest) Run(cell string) {
	if l := uint64(len(cell)); l > s.max {
		s.max = l
	}
	for i, class := range s.types {
		if class.failed {
			continue
		}
		if !class.regex.MatchString(cell) {
			s.types[i].failed = true
		}
	}
	return
}

// Passed always returns true
func (s *StringTest) Passed() bool {
	return s.max > 0
}

// PrintResult indicates either failure or the detected string class
func (s *StringTest) PrintResult(out io.Writer) {
	if !s.Passed() {
		fmt.Fprintf(out, stringResultFormat, "Fail", s.max)
		return
	}
	for _, class := range s.types {
		if class.failed {
			continue
		}
		fmt.Fprintf(out, stringResultFormat, class.name, s.max)
		return
	}
	fmt.Fprintf(out, stringResultFormat, "UTF-8", s.max)
	return
}
