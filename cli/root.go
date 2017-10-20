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
)

// GlobalFlags
type GlobalFlags struct {
    Delimiter  string `short:"d" long:"delimiter"   desc:"A delimiter to use instead of a comma"`
    RowData    bool   `short:"r" long:"row-data"    desc:"Data is in rows, not columns"`
	SkipLabels bool   `short:"s" long:"skip-labels" desc:"Skip the first row or column as labels"`
}

// Root is the main command for this application
var Root *cmd.RootCMD

func init() {
	// Build Application
	Root = &cmd.RootCMD{
		Name:  "csv-analyze",
		Short: "A useful tool for discovering the contents of a CSV file",
		Flags: &GlobalFlags{",",false,false},
	}
	// Setup the Sub-Commands
	Root.RegisterCMD(&cmd.Help)
	Root.RegisterCMD(&Empty)
	Root.RegisterCMD(&Types)
	Root.RegisterCMD(&Unique)
}