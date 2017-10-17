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
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/DataDrake/csv-analyze/tests"
	"io"
	"strings"
)

const columnFormat = "\n\033[1mResults for Column [%d]:\033[21m\n"

// Suite checks for empty cells
type Suite struct {
	groups    []tests.Group
	skipFirst bool
}

// NewSuite returns a fresh Suite
func NewSuite(skipFirst bool) tests.Suite {
	return &Suite{nil, skipFirst}
}

// Run reads an entire CSV, noting empty cells as it goes
func (t *Suite) Run(src *csv.Reader, dst io.Writer) (err error) {
	var row []string
	i := 0
	for err != io.EOF {
		row, err = src.Read()
		switch err {
		case nil:
			if i == 0 {
				t.groups = make([]tests.Group, len(row))
				for j := range row {
					t.groups[j] = NewGroup()
				}
				if t.skipFirst {
					i++
					continue
				}
			}
			for j, cell := range row {
				t.groups[j].Run(strings.TrimSpace(cell))
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
	for j, group := range t.groups {
		fmt.Fprintf(dst, columnFormat, j)
		group.PrintResult(dst)
	}
	return
}
