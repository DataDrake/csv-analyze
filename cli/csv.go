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
    "compress/bzip2"
    "compress/gzip"
    "encoding/csv"
	"fmt"
    "io"
	"os"
    "strings"
)

// OpenCSV opens a CSV file with buffered IO, handling bzip2 and gzip compression formats
func OpenCSV(filename string) (file *os.File, decompressor io.ReadCloser, reader *csv.Reader) {
    file, err := os.Open(filename)
    if err != nil {
        fmt.Printf("ERROR: failed to open file '%s', reason: %s\n", filename, err.Error())
        os.Exit(1)
    }
    buff := bufio.NewReader(file)
    switch {
    case strings.HasSuffix(filename, ".bz2"):
        bz := bzip2.NewReader(buff)
        reader = csv.NewReader(bz)
    case strings.HasSuffix(filename, ".gz"):
        decompressor, err = gzip.NewReader(buff)
        if err != nil {
            file.Close()
            fmt.Printf("ERROR: failed to open file '%s', reason: %s\n", filename, err.Error())
            os.Exit(1)
        }
        reader = csv.NewReader(decompressor)
    default:
        reader = csv.NewReader(buff)
    }
    return
}
