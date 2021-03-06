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

  - [x] Missing columns
  - [x] Empty cells
  - More TBD

### Data Types

  - [x] Integer
    - [x] uint8
    - [x] uint16
    - [x] uint32
    - [x] uint64
    - [x] int8
    - [x] int16
    - [x] int32
    - [x] int64
  - [x] Floating Point
    - [x] float32
    - [x] float64
  - [x] Timestamps
    - Everything in the `time` package plus an Apache log format
  - [x] Boolean
    - [x] Y/N
    - [x] yes/no
    - [x] true/false
    - [x] t/f
    - [x] 1/0
    - [x] 1.0/0.0
  - [x] Strings
    - [x] Alpha
    - [x] Alphanumeric
    - [x] ASCII
    - [x] UTF-8

### Uniqueness
In addition to the previously described functionality, `csv-analyze` 
also possesses the ability to tally the unique values found in a
specific column.

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
