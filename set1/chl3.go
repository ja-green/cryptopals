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

// Challenge 3 - Single-byte XOR cipher
//
// The hex encoded string:
// 1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736
// ... has been XOR'd against a single character. Find the key, decrypt the message.
//
// You can do this by hand. But don't: write code to do it for you.
//
// How? Devise some method for "scoring" a piece of English plaintext.
// Character frequency is a good metric. Evaluate each output and choose the one with the best score.

package set1

// This is an approximate character frequency array of the English language.
// It is used to determine which plaintext is closest to being English text. The
// table maps character to frequency, a number between 0 and 1 with each index referring
// to a character in alphabetical order and the last index referring to a space character.
// Note that this table implicitly assumes an ASCII representation of the english language.
var (
	index = 0
	index2 = 0
	CHAR_FREQ_EXP = [27]float64{
		0.0651738, 0.0124248, 0.0217339, 0.0349835,
		0.1041442, 0.0197881, 0.0158610, 0.0492888,
		0.0558094, 0.0009033, 0.0050529, 0.0331490,
		0.0202124, 0.0564513, 0.0596302, 0.0137645,
		0.0008606, 0.0497563, 0.0515760, 0.0729357,
		0.0225134, 0.0082903, 0.0171272, 0.0013692,
		0.0145984, 0.0007836, 0.1918182,
		}
)

func Challenge3(s string) string {
	min_score 	:= -1.0
	min_string 	:= ""

	for i := 0; i < 26; i++ {
		xorkey := make([]byte, len(s))
		for j := 0; j < len(xorkey); j++ {
			xorkey[j] = byte(97 + i)
		}

		hex_result 		:= Challenge2(string(HexEncode([]byte(s))), string(HexEncode(xorkey)))
		//hex_result 		:= Challenge2(strings.ToLower(s), string(HexEncode(xorkey)))
		plain_result 	:= string(HexDecode([]byte(hex_result)))

		chisq := ChiSq(plain_result)

		if chisq < min_score || min_score == -1 {
			min_score 	= chisq
			min_string 	= plain_result
		}
	}

	return min_string
}

func ChiSq(s string) float64 {
	char_cnt  	:= make([]int, 26)
	total_len	:= len(s)
	score		:= 0.0

	for i := 0; i < len(s); i++ {
		c := s[i]

		switch {
		case c >= 65 && c <= 90:	char_cnt[c - 65]++
		case c >= 97 && c <= 122: 	char_cnt[c - 97]++
		default:					total_len--
		}
	}

	for i := 0; i < 26; i++ {
		exp 	:= float64(total_len) * CHAR_FREQ_EXP[i]
		diff 	:= float64(char_cnt[i]) - exp
		score 	+= diff * diff / exp
	}

	return score
}
