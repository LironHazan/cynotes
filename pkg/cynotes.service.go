package cynotes

import (
	cryptoUtils "cynotes/pkg/crypto"
	fsutils "cynotes/pkg/fs"
	"cynotes/pkg/git"
	"encoding/hex"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func InitCYNotes(user string, repo string) error {
	// Greet
	uname, err := fsutils.GetUserName()
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("Hey %s\n", uname)

	path, err := fsutils.NormalizeCYNotesPath(uname)
	localRepoPath := path + "/" + repo
	fmt.Printf("%s \n", localRepoPath)
	if !fsutils.IsPathExists(localRepoPath) {
		// Create docs base folder
		err = os.Mkdir(path, 0755)
		if err != nil {
			log.Fatalf(err.Error())
		}
		if err != nil {
			return err
		} else {
			fmt.Printf("Cloning %s\n", repo)
			git.Clone(user, repo, path)
		}
	}
	return err
}

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

	visit := func(path string, dir fs.DirEntry, err error) error {
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
