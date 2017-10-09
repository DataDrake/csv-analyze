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

// Test represents one of the various checks performed
type Test interface {
	Run(string)
	Passed() bool
	PrintResult()
}

// Group is a set of related tests which may or may not depend o none another
type Group interface {
	Run(string)
	Advance()
	Done() bool
	PrintResults()
}
