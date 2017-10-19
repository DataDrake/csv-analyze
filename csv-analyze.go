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

package main

import (
	"github.com/DataDrake/cli-ng/cmd"
	"github.com/DataDrake/csv-analyze/cli"
)

func main() {

	// Global Flags
	flags := struct {
        RowData    bool `short:"r" long:"row-data"    desc:"Data is in rows, not columns"`
		SkipLabels bool `short:"s" long:"skip-labels" desc:"Skip the first row or column as labels"`
	}{}

	// Build Application
	r := &cmd.RootCMD{
		Name:  "csv-analyze",
		Short: "A useful tool for discovering the contents of a CSV file",
		Flags: &flags,
	}

	// Setup the Sub-Commands
	r.RegisterCMD(&cmd.Help)
	r.RegisterCMD(&cli.Empty)
	r.RegisterCMD(&cli.Types)
	r.RegisterCMD(&cli.Unique)

	// Run the program
	r.Run()
}
