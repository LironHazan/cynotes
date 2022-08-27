package cryptoUtils

import (
	"encoding/hex"
	"os"
	"testing"
)

func TestGetMD5Hash(t *testing.T) {
	hash := GetMD5Hash("../../test_files/use_me.txt")
	hashStr := hex.EncodeToString(hash[:])
	expected := "eb896d97585a06c9ab7e61e0c97615cf"
	if hashStr != expected {
		t.Errorf("got %s, expected %s", hashStr, expected)
	}
}

func TestEncryptDecryptFuncs(t *testing.T) {

	bytes, err := os.ReadFile("../../test_files/use_me.txt")
	if err != nil {
		t.Errorf("Cannot read test resource")
	}
	key := []byte("passphrase")
	ciphertext, err := EncryptAES(key, bytes)
	text, err := DecryptAES([]byte("passphrase"), ciphertext)
	decryptedTxt := string(text)
	expected := "Hey I'm a content which you want to encrypt"

	if decryptedTxt != expected {
		t.Errorf("got %s, expected %s", decryptedTxt, expected)
	}
}
