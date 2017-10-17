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
)

const groupFormat = "\t\033[96m\033[4m%s\033[0m\n"

// Group is a set of Type tests to run for a set of related values
type Group struct {
	tests map[string][]tests.Test
	names []string
}

// NewGroup creates a new test group for the type tests
func NewGroup() tests.Group {
	return &Group{
		map[string][]tests.Test{
			"Numerical": []tests.Test{
				NewUnsignedTest(),
				NewSignedTest(),
				NewFloatTest(),
			},
			"Logical": []tests.Test{
				NewBooleanTest(),
			},
			"DateTime": []tests.Test{
				NewTimeTest(),
			},
			"String": []tests.Test{
				NewStringTest(),
			},
		},
		[]string{"Numerical", "Logical", "DateTime", "String"},
	}
}

// Run hands the same string to all of the tests
func (g *Group) Run(cell string) {
	for _, ts := range g.tests {
		for _, t := range ts {
			t.Run(cell)
		}
	}
}

// PrintResult writes out the results of the type tests
func (g *Group) PrintResult(dst io.Writer) {
	for _, name := range g.names {
		fmt.Fprintf(dst, groupFormat, name)
		for _, t := range g.tests[name] {
			t.PrintResult(dst)
		}
	}
}
