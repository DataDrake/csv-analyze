# csv-analyze
Perform analytics on CSV files in order to better understand the kinds of data they represent

## Purpose
There is a treasure-trove of information currently available on the internet that is just
waiting for someone to come along and do something useful with it. Frequently this data will
find its way into a CSV format, but it may not have been properly sanitized for general use.
This program seeks to:

  * Perform analyses of CSV files to find non-uniformities in the contained data
  * Provide guidance on the kind of data in each column, in order to guide future
    reserialization of the contained information

### Checks

  - [ ] Missing columns
  - [ ] Empty cells
  - More TBD

### Data Types

  - [ ] Integer
    - [ ] uint8
    - [ ] uint16
    - [ ] uint32
    - [ ] uint64
    - [ ] int8
    - [ ] int16
    - [ ] int32
    - [ ] int64
  - [ ] Floating Point
    - [ ] float32
    - [ ] float64
  - [ ] Timestamps
    - Formats TBD, but probably everything in the `time` package plus an Apache log format
  - [ ] Boolean
    - [ ] Y/N
    - [ ] yes/no
    - [ ] true/false
  - [ ] Strings
    - [ ] Alpha
    - [ ] Alphanumeric
    - [ ] Extended

### License
Copyright 2017 Bryan T. Meyers <bmeyers@datadrake.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
