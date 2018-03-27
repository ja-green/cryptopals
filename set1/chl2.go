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

// Challenge 2 - Fixed XOR
//
// Write a function that takes two equal-length buffers and produces their XOR combination.
//
// If your function works properly, then when you feed it the string:
// 1c0111001f010100061a024b53535009181c
//
// ... after hex decoding, and when XOR'd against:
// 686974207468652062756c6c277320657965
//
// ... should produce:
// 746865206b696420646f6e277420706c6179

package set1

var (
	HEX_VAL = [16]byte{
		48, 49,  50,  51,
		52, 53,  54,  55,
		56, 57,  97,  98,
		99, 100, 101, 102,
	}
)

func Challenge2(s1 string, s2 string) string {
	b1 	:= HexDecode([]byte(s1))
	b2 	:= HexDecode([]byte(s2))
	n 	:= 0

	if len(b1) < len(b2) {
		n = len(b1)
	} else {
		n = len(b2)
	}

	res := make([]byte, n)

	for i := 0; i < n; i++ {
		res[i] = b1[i] ^ b2[i]
	}

	return string(HexEncode(res))
}

func HexEncode(src []byte) []byte {
	res := make([]byte, len(src) << 1)

	for i, v := range src {
		res[i << 1] 	= HEX_VAL[v >> 4]
		res[i << 1 + 1] = HEX_VAL[v & 0x0f]
	}

	return res
}
