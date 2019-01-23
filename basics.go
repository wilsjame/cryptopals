// cryptopals.com/sets/2
package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
)

func main() {
	fmt.Printf("Set 1\n")
	/**************************************/
	fmt.Printf("Convert hex to base64\n")
	const s = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	fmt.Printf("%v\n", hex2base64(s))
	/**************************************/
	fmt.Printf("Fixed XOR\n")
	const s2 = "1c0111001f010100061a024b53535009181c"
	const xor = "686974207468652062756c6c277320657965"
	fmt.Printf("%x\n", fixedXOR(s2, xor))
	/**************************************/
	fmt.Printf("Single-byte XOR cipher\n")
	const s3 = "1b37373331363f78151b7f2b783431334d78397828372d363c78373e783a393b3736"
	var char = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	var highscore = 0
	var score = 0

	for i, _ := range char {

		// reset score
		score = 0

		// copy cipher text into result
		result, err := hex.DecodeString(s3)
		if err != nil {
			log.Fatal(err)
		}

		// XOR result with the current character
		for j, _ := range result {
			result[j] = result[j] ^ char[i]
		}

		// score result based on letter frequency
		for i := 0; i < len(result); i++ {

			if result[i] == byte('e') || result[i] == byte('E') ||
				result[i] == byte('a') || result[i] == byte('A') ||
				result[i] == byte('t') || result[i] == byte('T') ||
				result[i] == byte('o') || result[i] == byte('O') ||
				result[i] == byte('n') || result[i] == byte('N') ||
				result[i] == byte('s') || result[i] == byte('S') ||
				result[i] == byte('h') || result[i] == byte('H') ||
				result[i] == byte('r') || result[i] == byte('R') ||
				result[i] == byte('d') || result[i] == byte('D') ||
				result[i] == byte('l') || result[i] == byte('L') ||
				result[i] == byte('u') || result[i] == byte('U') {
				score++
			}

		}

		if score > highscore {
			highscore = score
			fmt.Printf("%q : %v : %s\n", char[i], highscore, result)
		}

	}

}

/* Converts a hex string to base64
 * and returns string result. */
func hex2base64(s string) string {
	var raw []byte

	raw, err := hex.DecodeString(s)
	if err != nil {
		log.Fatal(err)
	}

	return base64.StdEncoding.EncodeToString(raw)

}

/* XOR's two equal-length hex strings
 * and returns the byte slice result */
func fixedXOR(s1, s2 string) []byte {
	var xor []byte
	var raw_1 []byte
	var raw_2 []byte

	raw_1, err := hex.DecodeString(s1)
	if err != nil {
		log.Fatal(err)
	}
	raw_2, err = hex.DecodeString(s2)
	if err != nil {
		log.Fatal(err)
	}

	for i, _ := range raw_1 {
		xor = append(xor, raw_1[i]^raw_2[i])
	}

	return xor
}
