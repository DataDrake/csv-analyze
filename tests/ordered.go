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

package test

// OrderedTestGroup is a set of tests which must be run sequentially
type OrderedTestGroup struct {
	tests   []Test
	current int
}

// Run execues the current test, unless it has already failed or there are no more tests
func (ordered *OrderedTestGroup) Run(cell string) {
	if ordered.Done() {
		return
	}
	ordered.tests[ordered.current].Run(cell)
	return
}

// Advance moves on to the next test, unless already finished
func (ordered *OrderedTestGroup) Advance() {
	if ordered.Done() {
		return
	}
	ordered.current++
	return
}

// Done checks if either all tests have been run or the current test has failed
func (ordered *OrderedTestGroup) Done() bool {
	if ordered.current >= len(ordered.tests) {
		return true
	}
	return !ordered.tests[ordered.current].Passed()
}

// PrintResults will print the result from the last test (if all passed) or the current failed test
func (ordered *OrderedtestGroup) PrintResults() {
	if l := len(ordered.tests); ordered.current >= l && l > 0 {
		ordered.tests[l-1].PrintResult()
		return
	}
	ordered.tests[ordered.current].PrintResult()
	return
}
