package cryptoUtils

import (
	"encoding/hex"
	"testing"
)

func TestGetMD5Hash(t *testing.T) {
	hash := GetMD5Hash("../../test_files/use_me.txt")
	hashStr := hex.EncodeToString(hash[:])
	expected := "01a93c748adffbbffac86eefea09f6c5"
	if hashStr != expected {
		t.Errorf("got %s, expected %s", hashStr, expected)
	}
}
