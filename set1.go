package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math"
	"strings"
)

var b64 = base64.StdEncoding

func hexToBase64(hexString string) string {
	str, err := hex.DecodeString(hexString)
	if err != nil {
		panic(err)
	}
	return b64.EncodeToString(str)
}

func xor(one, two []byte) []byte {
	answer := []byte{}
	for i, _ := range one {
		if i < len(two) {
			answer = append(answer, one[i]^two[i])
		}
	}
	return answer
}

// low numbers are better
func scoreEnglishText(input string) float64 {
	input = strings.ToLower(input)
	inputFrequency := map[rune]float64{}
	for _, c := range input {
		if val, ok := inputFrequency[c]; ok {
			inputFrequency[c] = val + float64(1)/float64(len(input))
		} else {
			inputFrequency[c] = float64(1) / float64(len(input))
		}
	}

	difference := float64(0)
	for char, freq := range letterFrequency {
		difference += math.Abs(freq - inputFrequency[char])
	}
	return difference
}

func main() {
	fmt.Print(hexToBase64("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"))
}
