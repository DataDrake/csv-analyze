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

import (
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/DataDrake/csv-analyze/tests"
	"io"
	"strings"
)

const uniqueColumnHeader = "\nUnique values for Column [%v]:\n"
const uniqueResultFormat = "\t%v: %v\n"

// Suite checks for unique values
type Suite struct {
	values []map[string]uint64
}

// NewSuite returns a fresh Suite
func NewSuite() tests.Suite {
	return &Suite{
		make([]map[string]uint64, 0),
	}
}

// Run reads an entire CSV, noting unique values as it goes
func (f *Suite) Run(src *csv.Reader, dst io.Writer) (err error) {
	var row []string
	i := 0
	for err != io.EOF {
		row, err = src.Read()
		switch err {
		case nil:
			if i == 0 {
				for j := 0; j < len(row); j++ {
					f.values = append(f.values, make(map[string]uint64))
				}
			}
			for j, cell := range row {
				cell = strings.TrimSpace(cell)
				f.values[j][cell]++
			}
			i++
		case io.EOF:
			continue
		default:
			err = errors.New("Failed to read next line, reason: " + err.Error())
			return
		}
	}
	err = nil
	for i, counts := range f.values {
		output := NewSorter()
		for k, v := range counts {
			output.Add(v, k)
		}
        output.Sort()
		fmt.Fprintf(dst, uniqueColumnHeader, i)
		for _, e := range output.values {
			fmt.Fprintf(dst, uniqueResultFormat, e.value, e.key)
		}
	}
	return
}
