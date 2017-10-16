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

package validation

import (
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/DataDrake/csv-analyze/tests"
	"io"
	"strings"
)

const emptyResultFormat = "\tRow: %v, Col: %v\n"

// EmptySuite checks for empty cells
type EmptySuite struct{}

// NewEmptySuite returns a fresh EmptySuite
func NewEmptySuite() tests.Suite {
	return &EmptySuite{}
}

// Run reads an entire CSV, noting empty cells as it goes
func (f *EmptySuite) Run(src *csv.Reader, dst io.Writer) (err error) {
	var row []string
	i := 0
	for {
		row, err = src.Read()
		switch err {
		case nil:
			for j, cell := range row {
				if len(strings.TrimSpace(cell)) == 0 {
					fmt.Fprintf(dst, emptyResultFormat, i, j)
				}
			}
			i++
		case io.EOF:
			err = nil
			return
		default:
			err = errors.New("Failed to read next line, reason: " + err.Error())
			return
		}
	}
}
