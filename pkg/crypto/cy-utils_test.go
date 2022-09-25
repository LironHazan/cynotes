package cryptoUtils

import (
	"encoding/hex"
	_ "github.com/samber/mo"
	"os"
	"testing"
)

var fpcy = FpCyUtil{}

func TestGetMD5Hash(t *testing.T) {
	hash := fpcy.GetMD5Hash("../../test_files/use_me.txt").OrEmpty()
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
	ciphertext := fpcy.EncryptAES(key, bytes).RightOrEmpty()
	text, _ := fpcy.DecryptAES([]byte("passphrase"), ciphertext).Right()
	decryptedTxt := string(text)
	expected := "Hey I'm a content which you want to encrypt"

	if decryptedTxt != expected {
		t.Errorf("got %s, expected %s", decryptedTxt, expected)
	}
}
