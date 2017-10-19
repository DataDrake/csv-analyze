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
		make([]*pair,0),
	}
}

// Insert adds a new item to the Sorter, using an insertion sort
func (s *Sorter) Insert(key uint64, value string) (err error) {
    p := &pair{key,value}
    for i := 0; i < len(s.values); i++ {
        if p.key >= s.values[i].key {
            temp := s.values[i]
            s.values[i] = p
            p = temp
        }
    }
    s.values = append(s.values, p)
	return
}
