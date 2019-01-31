// cryptopals.com
package main

import (
	"bufio"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"math/bits"
	"os"
)

var s string

func main() {
	fmt.Printf("Set 1\n")

	fmt.Printf("/**************************************/\n")
	fmt.Printf("Convert hex to base64\n")
	s = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	fmt.Printf("%v\n", hex2base64(s))

	fmt.Printf("/**************************************/\n")
	fmt.Printf("Fixed XOR\n")
	s = "1c0111001f010100061a024b53535009181c"
	const xor = "686974207468652062756c6c277320657965"
	fmt.Printf("%x\n", fixedXOR(s, xor))

	fmt.Printf("/**************************************/\n")
	fmt.Printf("Single-byte XOR cipher\n")
	s = "1b37373331363f78151b7f2b783431334d78397828372d363c78373e783a393b3736"
	score, key, message := byteXORcipher(s)
	fmt.Printf("%v : %q : %+q\n", score, key, message)

	fmt.Printf("/**************************************/\n")
	fmt.Printf("Detect single-character XOR\n")
	f, err := os.Open("4.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)

	// loop over all line in the file
	for scanner.Scan() {
		//TODO show highest scores first
		score, key, message := byteXORcipher(scanner.Text())
		fmt.Printf("%v : %q : %+q\n", score, key, message)
	}

	fmt.Printf("/**************************************/\n")
	fmt.Printf("Implementing repeating-key XOR\n")
	s = "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	sol := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"

	if hex.EncodeToString(XORencrypt("ICE", s)) == sol {
		fmt.Printf("equal!\n")
	} else {
		fmt.Printf("no match :(\n")
	}

	fmt.Printf("/**************************************/\n")
	fmt.Printf("Break repeating-key XOR\n")
	fmt.Printf("Hamming distance %v\n", hammingD("this is a test", "wokka wokka!!!"))

}

/* Computes the Hamming distance between
 * two strings (i.e. the number of differing
 * bits). */
func hammingD(s1, s2 string) int {
	var hd = 0
	var b1 = []byte(s1)
	var b2 = []byte(s2)

	for i, _ := range b1 {
		// XOR and count set bits
		hd += bits.OnesCount(uint(b1[i] ^ b2[i]))
	}

	return hd

}

/* Encrypts a string using repeating-key XOR.
 * Returns byte slice result. */
func XORencrypt(key, s string) []byte {
	k := []byte(key) // key
	t := []byte(s)   // plaintext
	var ct []byte    // ciphertext

	for i, _ := range t {
		ct = append(ct, k[i%len(k)]^t[i])
	}

	return ct

}

/* Takes a hex string that has been XOR'd
 * against a single character (i.e. key).
 * Finds the key by scoring the English plaintext
 * using character frequency as a metric.
 * Returns the score, key, and decrpyted message. */
func byteXORcipher(s string) (int, byte, []byte) {
	var char = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789") // add more keys if necessary
	var message []byte
	var key byte
	var highscore = 0

	for i, _ := range char {

		// reset score
		score := 0

		// copy cipher text into result
		result, err := hex.DecodeString(s)
		if err != nil {
			log.Fatal(err)
		}

		// XOR result with the current character
		for j, _ := range result {
			result[j] = result[j] ^ char[i]
		}

		// score result based on letter frequency
		for i := 0; i < len(result); i++ {
			letter := result[i]

			// EATON SHRDLU
			if letter == byte('e') || letter == byte('E') ||
				letter == byte('a') || letter == byte('A') ||
				letter == byte('t') || letter == byte('T') ||
				letter == byte('o') || letter == byte('O') ||
				letter == byte('n') || letter == byte('N') ||
				letter == byte('s') || letter == byte('S') ||
				letter == byte('h') || letter == byte('H') ||
				letter == byte('r') || letter == byte('R') ||
				letter == byte('d') || letter == byte('D') ||
				letter == byte('l') || letter == byte('L') ||
				letter == byte('u') || letter == byte('U') {
				score++
			}

			// TH HE AN RE ER IN ON AT ND ST ES EN OF TE ED OR TI HI AS TO
			// LL EE SS OO TT FF RR NN PP CC
			if i < len(result)-1 {
				nextLetter := result[i+1]

				if letter == byte('t') || letter == byte('T') {

					if nextLetter == byte('h') || nextLetter == byte('H') ||
						nextLetter == byte('e') || nextLetter == byte('E') ||
						nextLetter == byte('i') || nextLetter == byte('I') ||
						nextLetter == byte('o') || nextLetter == byte('O') ||
						nextLetter == byte('t') || nextLetter == byte('T') {
						score++
					}

				}

				if letter == byte('h') || letter == byte('H') {

					if nextLetter == byte('e') || nextLetter == byte('E') ||
						nextLetter == byte('i') || nextLetter == byte('I') {
						score++
					}

				}

				if letter == byte('a') || letter == byte('A') {

					if nextLetter == byte('n') || nextLetter == byte('N') ||
						nextLetter == byte('t') || nextLetter == byte('T') ||
						nextLetter == byte('s') || nextLetter == byte('S') {
						score++
					}

				}

				if letter == byte('r') || letter == byte('R') {

					if nextLetter == byte('e') || nextLetter == byte('E') ||
						nextLetter == byte('r') || nextLetter == byte('R') {
						score++
					}

				}

				if letter == byte('e') || letter == byte('E') {

					if nextLetter == byte('r') || nextLetter == byte('R') ||
						nextLetter == byte('s') || nextLetter == byte('S') ||
						nextLetter == byte('n') || nextLetter == byte('N') ||
						nextLetter == byte('d') || nextLetter == byte('D') ||
						nextLetter == byte('e') || nextLetter == byte('E') {
						score++
					}

				}

				if letter == byte('i') || letter == byte('I') {

					if nextLetter == byte('n') || nextLetter == byte('N') {
						score++
					}

				}

				if letter == byte('o') || letter == byte('O') {

					if nextLetter == byte('n') || nextLetter == byte('N') ||
						nextLetter == byte('f') || nextLetter == byte('F') ||
						nextLetter == byte('r') || nextLetter == byte('R') ||
						nextLetter == byte('o') || nextLetter == byte('O') {
						score++
					}

				}

				if letter == byte('n') || letter == byte('N') {

					if nextLetter == byte('d') || nextLetter == byte('D') ||
						nextLetter == byte('n') || nextLetter == byte('N') {
						score++
					}

				}

				if letter == byte('s') || letter == byte('S') {

					if nextLetter == byte('t') || nextLetter == byte('T') ||
						nextLetter == byte('s') || nextLetter == byte('S') {
						score++
					}

				}

				if letter == byte('l') || letter == byte('L') {

					if nextLetter == byte('l') || nextLetter == byte('L') {
						score++
					}

				}

				if letter == byte('f') || letter == byte('F') {

					if nextLetter == byte('f') || nextLetter == byte('F') {
						score++
					}

				}

				if letter == byte('p') || letter == byte('P') {

					if nextLetter == byte('p') || nextLetter == byte('P') {
						score++
					}

				}

				if letter == byte('c') || letter == byte('C') {

					if nextLetter == byte('c') || nextLetter == byte('C') {
						score++
					}

				}

			}

		}

		if score > highscore {
			highscore = score
			key = char[i]
			message = result
			// debug
			//fmt.Printf(" %v : %q : %+q\n", highscore, key, message)
		}

	}

	return highscore, key, message

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
 * and returns the byte slice result. */
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
