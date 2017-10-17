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
    "github.com/DataDrake/csv-analyze/tests/types"
	"os"
)

// Types reads the CSV file and determines the possible types
var Types = cmd.CMD{
	Name:  "types",
	Alias: "T",
	Short: "Try to determine the possible data types",
	Args:  TypesArgs,
	Run:   TypesRun,
}

// TypesArgs contains the arguments for the "empty" subcommand
var TypesArgs = struct {
	csv string `desc:"Path to a CSV file to analyze"`
}{}

// TypesRun carries out the type testing process
func TypesRun(r *cmd.RootCMD, c *cmd.CMD) {
	fn := flag.Arg(1)
    csvFile, err := os.Open(fn)
    if err != nil {
        fmt.Printf("ERROR: failed to open file '%s', reason: %s\n", fn, err.Error())
        os.Exit(1)
    }
    defer csvFile.Close()
    buff := bufio.NewReader(csvFile)
    reader := csv.NewReader(buff)
    suite := types.NewSuite(false)
    suite.Run(reader, os.Stdout)
}
