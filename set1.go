package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
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

func main() {
	fmt.Print(hexToBase64("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"))
}
