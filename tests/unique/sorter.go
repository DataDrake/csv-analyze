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

package unique

type pair struct {
	key   uint64
	value string
}

// Sorter sorts a set of key-value pairs by key
type Sorter struct {
	values []*pair
}

// NewSorter returns a fresh Sorter
func NewSorter() *Sorter {
	return &Sorter{
		make([]*pair, 0),
	}
}

// Add stores a new item in the Sorter
func (s *Sorter) Add(key uint64, value string) {
	s.values = append(s.values, &pair{key, value})
	return
}
