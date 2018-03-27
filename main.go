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

package main

import (
	"os"
	"fmt"
	"strconv"

	"github.com/ja-green/cryptopals/set1"
	"github.com/ja-green/cryptopals/data"
)

const (
	SET_1 = iota + 1
	SET_2
	SET_3
	SET_4
	SET_5
	SET_6
	SET_7
	SET_8
)

const (
	CHL_1 = iota + 1
	CHL_2
	CHL_3
	CHL_4
	CHL_5
	CHL_6
	CHL_7
	CHL_8
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "[fatal] two arguments must be given\n")
		os.Exit(1)
	}

	set, set_err := strconv.Atoi(os.Args[1])
	chl, chl_err := strconv.Atoi(os.Args[2])

	if set_err != nil || chl_err != nil {
		fmt.Fprintf(os.Stderr, "[fatal] arguments must be numbers\n")
		os.Exit(1)
	}

	switch set {
	case SET_1:	switch chl {
		case CHL_1:	fmt.Println(set1.Challenge1(data.SET_1_CHL_1))
		case CHL_2:	fmt.Println(set1.Challenge2(data.SET_1_CHL_2_1, data.SET_1_CHL_2_2))
		case CHL_3: fmt.Println(set1.Challenge3(data.SET_1_CHL_3))
		case CHL_4: fmt.Println(set1.Challenge4(data.SET_1_CHL_4))
		case CHL_5:
		case CHL_6:
		case CHL_7:
		case CHL_8:
		default: 	fmt.Fprintf(os.Stderr, "[fatal] unknown challenge '%d'\n", chl)
	}

	case SET_2:	switch chl {
		case CHL_1:
		case CHL_2:
		case CHL_3:
		case CHL_4:
		case CHL_5:
		case CHL_6:
		case CHL_7:
		case CHL_8:
		default: 	fmt.Fprintf(os.Stderr, "[fatal] unknown challenge '%d'\n", chl)
	}

	case SET_3:	switch chl {
		case CHL_1:
		case CHL_2:
		case CHL_3:
		case CHL_4:
		case CHL_5:
		case CHL_6:
		case CHL_7:
		case CHL_8:
		default: 	fmt.Fprintf(os.Stderr, "[fatal] unknown challenge '%d'\n", chl)
	}

	case SET_4:	switch chl {
		case CHL_1:
		case CHL_2:
		case CHL_3:
		case CHL_4:
		case CHL_5:
		case CHL_6:
		case CHL_7:
		case CHL_8:
		default: 	fmt.Fprintf(os.Stderr, "[fatal] unknown challenge '%d'\n", chl)
	}

	case SET_5:	switch chl {
		case CHL_1:
		case CHL_2:
		case CHL_3:
		case CHL_4:
		case CHL_5:
		case CHL_6:
		case CHL_7:
		case CHL_8:
		default: 	fmt.Fprintf(os.Stderr, "[fatal] unknown challenge '%d'\n", chl)
	}

	case SET_6:	switch chl {
		case CHL_1:
		case CHL_2:
		case CHL_3:
		case CHL_4:
		case CHL_5:
		case CHL_6:
		case CHL_7:
		case CHL_8:
		default: 	fmt.Fprintf(os.Stderr, "[fatal] unknown challenge '%d'\n", chl)
	}

	case SET_7:	switch chl {
		case CHL_1:
		case CHL_2:
		case CHL_3:
		case CHL_4:
		case CHL_5:
		case CHL_6:
		case CHL_7:
		case CHL_8:
		default: 	fmt.Fprintf(os.Stderr, "[fatal] unknown challenge '%d'\n", chl)
	}

	case SET_8:	switch chl {
		case CHL_1:
		case CHL_2:
		case CHL_3:
		case CHL_4:
		case CHL_5:
		case CHL_6:
		case CHL_7:
		case CHL_8:
		default: 	fmt.Fprintf(os.Stderr, "[fatal] unknown challenge '%d'\n", chl)

	}

	default:		fmt.Fprintf(os.Stderr, "[fatal] unknown set '%d'\n", set)
	}
}
