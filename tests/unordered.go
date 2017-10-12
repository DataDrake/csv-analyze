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

package tests

import "io"

// UnorderedTestGroup is a set of tests which may be run all at once
type UnorderedTestGroup struct {
	tests   []Test
	current int
}

// Run executes the current test, if not already done
func (unordered *UnorderedTestGroup) Run(cell string) {
	if unordered.Done() {
		return
	}
	unordered.tests[unordered.current].Run(cell)
	return
}

// Advance moves on to the next test in the suite if not already done
func (unordered *UnorderedTestGroup) Advance() {
	if unordered.Done() {
		return
	}
	unordered.current++
}

// Done checks if either all tests have been run
func (unordered *UnorderedTestGroup) Done() bool {
	return unordered.current >= len(unordered.tests)
}

// PrintResults will print the result for each of the tests
func (unordered *UnorderedTestGroup) PrintResults(out io.Writer) {
	for _, t := range unordered.tests {
		t.PrintResult(out)
	}
	return
}
