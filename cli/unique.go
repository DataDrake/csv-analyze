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
    "github.com/DataDrake/cli-ng/cmd"
    "github.com/DataDrake/csv-analyze/tests/unique"
	"os"
)

// Unique checks for unique values in a CSV file
var Unique = cmd.CMD{
	Name:  "unique",
	Alias: "U",
	Short: "Check for unique values in a CSV file",
	Args:  &UniqueArgs{},
	Run:   UniqueRun,
}

// UniqueArgs contains the arguments for the "unique" subcommand
type UniqueArgs struct {
	CSV string `desc:"Path to a CSV file to analyze"`
}

// UniqueRun carries out the search unique CSV values
func UniqueRun(r *cmd.RootCMD, c *cmd.CMD) {
	args := c.Args.(*UniqueArgs)
    file, decompressor, reader := OpenCSV(args.CSV, r.Flags.(*GlobalFlags).Delimiter)
    if decompressor != nil {
        defer decompressor.Close()
    }
    defer file.Close()
    suite := unique.NewSuite()
    err := suite.Run(reader, os.Stdout)
    if err != nil {
        println(err.Error())
        os.Exit(1)
    }
}
