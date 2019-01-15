// cryptopals.com/sets/1
package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
)

func main() {
	print("Set 1\n")
	/**************************************/
	print("Convert hex to base64\n")
	const s = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	decoded, err := hex.DecodeString(s)
	if err != nil {
		log.Fatal(err)
	}
	sol_1 := base64.StdEncoding.EncodeToString(decoded)
	fmt.Printf("%v\n", sol_1)
	/**************************************/
	print("Fixed XOR\n")
	const s2 = "1c0111001f010100061a024b53535009181c"
	decoded_s2, err := hex.DecodeString(s2)
	if err != nil {
		log.Fatal(err)
	}
	const xor = "686974207468652062756c6c277320657965"
	decoded_xor, err := hex.DecodeString(xor)
	if err != nil {
		log.Fatal(err)
	}
	var fixed_xor []byte
	for i, _ := range decoded_s2 {
		fixed_xor = append(fixed_xor, decoded_s2[i]^decoded_xor[i])
	}
	sol_2 := hex.EncodeToString(fixed_xor)
	fmt.Printf("%s\n", sol_2)
	/**************************************/

}
