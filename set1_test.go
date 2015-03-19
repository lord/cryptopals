package main

import (
	"encoding/hex"
	"testing"
)

func TestCh1(t *testing.T) {
	str := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	answer := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	if hexToBase64(str) != answer {
		t.Error("Challenge 1 Failed")
	}
}

func TestCh2(t *testing.T) {
	str1, _ := hex.DecodeString("1c0111001f010100061a024b53535009181c")
	str2, _ := hex.DecodeString("686974207468652062756c6c277320657965")
	if hex.EncodeToString(xor(str1, str2)) != "746865206b696420646f6e277420706c6179" {
		t.Error("Challenge 2 Failed")
	}
}

func TestCh3(t *testing.T) {
	cryptoText, _ := hex.DecodeString("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	_, key := findSingleCharacterXorPlaintext(cryptoText)
	if key != 88 && key != 120 {
		t.Error("Challenge 3 Failed")
	}
}
