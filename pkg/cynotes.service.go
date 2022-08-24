package cynotes

import (
	cryptoUtils "cynotes/pkg/crypto"
	fsutils "cynotes/pkg/fs"
	"encoding/hex"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
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

func List() {
	// todo: model: notesMap := make(map[string][]string)

	visit := func(path string, di fs.DirEntry, err error) error {
		fmt.Println(path)
		return nil
	}

	path, _ := fsutils.GetCYNotesPath()
	err := filepath.WalkDir(path, visit)
	if err != nil {
		fmt.Println(err)
	}
}

// todo - implement view file
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
