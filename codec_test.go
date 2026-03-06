package builtin

import (
	"testing"

	"github.com/infrago/infra"
)

func TestDigitHelpersRoundTrip(t *testing.T) {
	encoded, err := EncryptDIGIT(123456)
	if err != nil {
		t.Fatalf("encrypt digit: %v", err)
	}

	decoded, err := DecryptDIGIT(encoded)
	if err != nil {
		t.Fatalf("decrypt digit: %v", err)
	}
	if decoded != 123456 {
		t.Fatalf("unexpected digit round trip: %d", decoded)
	}

	itemsEncoded, err := EncryptDIGITS([]int64{3, 5, 8})
	if err != nil {
		t.Fatalf("encrypt digits: %v", err)
	}

	itemsDecoded, err := DecryptDIGITS(itemsEncoded)
	if err != nil {
		t.Fatalf("decrypt digits: %v", err)
	}
	if len(itemsDecoded) != 3 || itemsDecoded[0] != 3 || itemsDecoded[1] != 5 || itemsDecoded[2] != 8 {
		t.Fatalf("unexpected digits round trip: %#v", itemsDecoded)
	}
}

func TestTextHelpersRoundTrip(t *testing.T) {
	encoded, err := EncryptTEXT("hello-world")
	if err != nil {
		t.Fatalf("encrypt text: %v", err)
	}

	decoded, err := DecryptTEXT(encoded)
	if err != nil {
		t.Fatalf("decrypt text: %v", err)
	}
	if decoded != "hello-world" {
		t.Fatalf("unexpected text round trip: %q", decoded)
	}

	itemsEncoded, err := EncryptTEXTS([]string{"alpha", "beta", "gamma"})
	if err != nil {
		t.Fatalf("encrypt texts: %v", err)
	}

	itemsDecoded, err := DecryptTEXTS(itemsEncoded)
	if err != nil {
		t.Fatalf("decrypt texts: %v", err)
	}
	if len(itemsDecoded) != 3 || itemsDecoded[0] != "alpha" || itemsDecoded[1] != "beta" || itemsDecoded[2] != "gamma" {
		t.Fatalf("unexpected texts round trip: %#v", itemsDecoded)
	}
}

func TestBuiltinRegistersDefaultCodecs(t *testing.T) {
	builtin.Setup()

	encodedText, err := infra.Encrypt(TEXT, "hello")
	if err != nil {
		t.Fatalf("encrypt builtin text codec: %v", err)
	}
	decodedText, err := infra.Decrypt(TEXT, encodedText)
	if err != nil {
		t.Fatalf("decrypt builtin text codec: %v", err)
	}
	if string(decodedText.([]byte)) != "hello" {
		t.Fatalf("unexpected builtin text decode: %#v", decodedText)
	}

	encodedTexts, err := infra.Encrypt(TEXTS, []string{"a", "b"})
	if err != nil {
		t.Fatalf("encrypt builtin texts codec: %v", err)
	}
	decodedTexts, err := infra.Decrypt(TEXTS, encodedTexts)
	if err != nil {
		t.Fatalf("decrypt builtin texts codec: %v", err)
	}
	items, ok := decodedTexts.([]string)
	if !ok || len(items) != 2 || items[0] != "a" || items[1] != "b" {
		t.Fatalf("unexpected builtin texts decode: %#v", decodedTexts)
	}

	encodedDigit, err := infra.Encrypt(DIGIT, int64(42))
	if err != nil {
		t.Fatalf("encrypt builtin digit codec: %v", err)
	}
	decodedDigit, err := infra.Decrypt(DIGIT, encodedDigit)
	if err != nil {
		t.Fatalf("decrypt builtin digit codec: %v", err)
	}
	if decodedDigit.(int64) != 42 {
		t.Fatalf("unexpected builtin digit decode: %#v", decodedDigit)
	}

	encodedDigits, err := infra.Encrypt(DIGITS, []int64{2, 4})
	if err != nil {
		t.Fatalf("encrypt builtin digits codec: %v", err)
	}
	decodedDigits, err := infra.Decrypt(DIGITS, encodedDigits)
	if err != nil {
		t.Fatalf("decrypt builtin digits codec: %v", err)
	}
	digits, ok := decodedDigits.([]int64)
	if !ok || len(digits) != 2 || digits[0] != 2 || digits[1] != 4 {
		t.Fatalf("unexpected builtin digits decode: %#v", decodedDigits)
	}
}
