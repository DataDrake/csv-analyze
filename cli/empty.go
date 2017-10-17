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

package cli

import (
    "bufio"
    "encoding/csv"
	"flag"
	"fmt"
    "github.com/DataDrake/cli-ng/cmd"
    "github.com/DataDrake/csv-analyze/tests/validation"
	"os"
)

// Empty checks for empty cells in a CSV file
var Empty = cmd.CMD{
	Name:  "empty",
	Alias: "E",
	Short: "Check for empty cells in a CSV file",
	Args:  EmptyArgs,
	Run:   EmptyRun,
}

// EmptyArgs contains the arguments for the "empty" subcommand
var EmptyArgs = struct {
	csv string `desc:"Path to a CSV file to analyze"`
}{}

// EmptyRun carries out the search for empty CSV cells
func EmptyRun(r *cmd.RootCMD, c *cmd.CMD) {
	fn := flag.Arg(1)
    csvFile, err := os.Open(fn)
    if err != nil {
        fmt.Printf("ERROR: failed to open file '%s', reason: %s\n", fn, err.Error())
        os.Exit(1)
    }
    defer csvFile.Close()
    buff := bufio.NewReader(csvFile)
    reader := csv.NewReader(buff)
    suite := validation.NewEmptySuite()
    suite.Run(reader, os.Stdout)
}
