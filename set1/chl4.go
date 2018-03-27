// Copyright 2018 Jack Green (ja-green)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Challenge 4 - Detect single-character XOR
// One of the 60-character strings in this file has been encrypted by single-character XOR.
//
// Find it.

package set1

import (
	"os"
	"fmt"
	"bufio"
)

func Challenge4(fname string) string {
	file, err := os.Open(fname)

	if err != nil {
		fmt.Fprintf(os.Stderr, "[fatal] file '%s' not found\n", fname)
		os.Exit(1)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "[fatal] error reading file '%s'\n", fname)
		os.Exit(1)
	}

	return ""
}
