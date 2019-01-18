// cryptopals.com/sets/1
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
	const s3 = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	var char = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")

	// XOR each character
	for i, _ := range char {
		// copy original string into result
		result, err := hex.DecodeString(s3)
		if err != nil {
			log.Fatal(err)
		}

		// XOR result with single character
		for j, _ := range result {
			result[j] = result[j] ^ char[i]
		}
		fmt.Printf("%+q\n", result)
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
