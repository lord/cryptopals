package main

import (
	"encoding/hex"
	"testing"
)

func TestCh1(t *testing.T) {
	str := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	answer := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	if hexToBase64(str) != answer {
		t.Error("Challange 1 Failed")
	}
}

func TestCh2(t *testing.T) {
	str1, _ := hex.DecodeString("1c0111001f010100061a024b53535009181c")
	str2, _ := hex.DecodeString("686974207468652062756c6c277320657965")
	answer, _ := hex.DecodeString("746865206b696420646f6e277420706c6179")
	if xor(str1, str2) != answer {
		t.Error("Challange 2 Failed")
	}
}
