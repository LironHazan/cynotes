package cynotes

import (
	cryptoUtils "cynotes/pkg/crypto"
	fsutils "cynotes/pkg/fs"
	"encoding/hex"
	"fmt"
	"os"
)

func Commit(filepath string, passphrase string) {
	filename := fsutils.ExtractFilename(filepath)
	basePath := fsutils.CreateNoteFolder(filename)

	hash := cryptoUtils.GetMD5Hash(filepath)
	hashStr := hex.EncodeToString(hash[:])
	commit := basePath + "/" + hashStr

	bytes, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Printf("Failed reading data from file: %s", err)
	}

	key := []byte(passphrase)
	ciphertext, err := cryptoUtils.EncryptAES(key, bytes)
	fmt.Printf("key %s \n", key)

	err = os.WriteFile(commit, ciphertext, 0755)
	if err != nil { // todo dont panic
		fmt.Printf("Failed reading data from file: %s", err)
	}

}

//func Run(passphrase string, filename string) error {
//	bytes, err := os.ReadFile(filepath)
//	if err != nil {
//		fmt.Printf("Failed reading data from file: %s", err)
//	}
//
//	text, err := cryptoUtils.DecryptAES([]byte(passphrase), ciphertext)
//	s := string(text)
//	fmt.Printf("s %s\n ", s)
//	return err
//}
