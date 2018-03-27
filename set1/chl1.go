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

// Challenge 1 - Convert hex to base64:
//
// The string:
// 49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d
//
// Should produce:
// SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t

package set1

// This is an array containing all the possible characters in a base64 encoded text.
// It is used to populate the output array with the correct characters. Each character is
// represented by its decimal ASCII value.
var (
	B64_VAL = [64]byte{
		65,  66,  67,  68,  69,  70,  71,  72,
		73,  74,  75,  76,  77,  78,  79,  80,
		81,  82,  83,  84,  85,  86,  87,  88,
		89,  90,  97,  98,  99,  100, 101, 102,
		103, 104, 105, 106, 107, 108, 109, 110,
		111, 112, 113, 114, 115, 116, 117, 118,
		119, 120, 121, 122, 48,  49,  50,  51,
		52,  53,  54,  55,  56,  57,  43,  47,
	}
)

// Challenge1 converts the given string, s into its representative byte array and
// calls the B64Encode function, returning a string representation of the resultant byte array.
//
// This function is meant as a wrapper to allow for more readable code in main.go
func Challenge1(s string) string {
	src := HexDecode([]byte(s))
	res := B64Encode(src)

	return string(res)
}

func B64Encode(src []byte) []byte {
	res := make([]byte, (len(src) + 2) / 3 * 4)

	res_idx, src_idx := 0, 0
	for src_idx < ((len(src) / 3) * 3) {
		x := uint(src[src_idx + 0]) << 16 | uint(src[src_idx + 1]) << 8 | uint(src[src_idx + 2])

		res[res_idx + 0] = B64_VAL[x >> 18 & 0x3F]
		res[res_idx + 1] = B64_VAL[x >> 12 & 0x3F]
		res[res_idx + 2] = B64_VAL[x >> 6  & 0x3F]
		res[res_idx + 3] = B64_VAL[x >> 0  & 0x3F]

		src_idx += 3
		res_idx += 4
	}

	switch len(src) - src_idx {
	case 1:
		x := uint(src[src_idx + 0]) << 16
		res[res_idx + 0] = B64_VAL[x >> 18 & 0x3F]
		res[res_idx + 1] = B64_VAL[x >> 12 & 0x3F]
		res[res_idx + 2] = 61
		res[res_idx + 3] = 61
	case 2:
		x := uint(src[src_idx + 0]) << 16 | uint(src[src_idx + 1]) << 8
		res[res_idx + 0] = B64_VAL[x >> 18 & 0x3F]
		res[res_idx + 1] = B64_VAL[x >> 12 & 0x3F]
		res[res_idx + 2] = B64_VAL[x >> 6  & 0x3F]
		res[res_idx + 3] = 61

	}

	return res

}

func HexDecode(src []byte) []byte {
	res := make([]byte, len(src) >> 1)

	for i := 0; i < len(src) >> 1; i++ {
		x := HexToByte(src[i << 1])
		y := HexToByte(src[i << 1 + 1])

		res[i] = (x << 4) | y
	}

	return res
}

func HexToByte(b byte) byte {
	switch {
	case '0' <= b && b <= '9':	return b - 48
	case 'a' <= b && b <= 'f':	return b - 'a' + 10
	case 'A' <= b && b <= 'F':	return b - 'A' + 10
	default:				  	return 0
	}
}
