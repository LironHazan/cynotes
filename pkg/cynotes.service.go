package cynotes

import (
	cryptoUtils "cynotes/pkg/crypto"
	"cynotes/pkg/editor"
	fsutils "cynotes/pkg/fs"
	"cynotes/pkg/git"
	promptUiUtils "cynotes/pkg/ui"
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
		return err
	}
	fmt.Printf("Hey %s\n", uname)

	path, err := fsutils.NormalizeCYNotesPath(uname)
	localRepoPath := path + "/" + repo
	fmt.Printf("%s \n", localRepoPath)
	if !fsutils.IsPathExists(localRepoPath) {
		// Create docs base folder
		err = os.Mkdir(path, 0755)
		if err != nil {
			return err
		}
		fmt.Printf("Cloning %s\n", repo)
		err := git.Clone(user, repo, path)
		if err != nil {
			return err
		}
		fsutils.CreateInitFile(repo)
	}
	return nil
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

func New(name string) error {
	notesDir := fsutils.GetWorkingRepoDir()

	// Create new note folder under the repo
	err := os.Mkdir(notesDir+"/"+name, 0755)
	if err != nil {
		log.Printf("could not create folder")
		return err
	}

	// tmp file
	tmpFilePath := notesDir + "/" + name + "/tmp"
	_, err = os.Create(tmpFilePath)
	if err != nil {
		log.Printf("could not create tmp file")
		return err
	}

	// open file for editing
	err = editor.ViEdit(tmpFilePath)
	if err != nil {
		return err
	}

	// rename to hash
	hash := cryptoUtils.GetMD5Hash(tmpFilePath)
	hashStr := hex.EncodeToString(hash[:])
	secretNote := notesDir + "/" + name + "/" + hashStr
	err = os.Rename(tmpFilePath, secretNote)
	if err != nil {
		log.Printf("could not rename tmp file")
		return err
	}

	// Encrypt
	log.Printf("Enter passphrase for encrypting note")
	passphrase, _ := promptUiUtils.PromptPasswdInput()

	bytes, err := os.ReadFile(secretNote)
	if err != nil {
		log.Printf("Failed reading data from file: %s", err)
		return err
	}

	key := []byte(passphrase)
	ciphertext, err := cryptoUtils.EncryptAES(key, bytes)
	fmt.Printf("key %s \n", key)

	err = os.WriteFile(secretNote, ciphertext, 0755)
	if err != nil {
		log.Printf("Failed reading data from file: %s", err)
		return err
	}

	selection, _ := promptUiUtils.BasicPromptSelections("Do you want to push changes?", []string{"Yes", "No"})
	if selection == "Yes" {
		git.Add(notesDir, secretNote)
		git.Commit(notesDir)
		git.Push(notesDir)
	}
	return nil
}

func EditNote(name string) {
	// fetch latest revision
	// copy content to tmpFile
	// edit tmpFile
	// save + push new revision
	notesDir := fsutils.GetWorkingRepoDir()
	noteDir := notesDir + "/" + name
	log.Printf(noteDir)
	var maxModTime int64 = 0
	note := ""
	visit := func(path string, dir fs.DirEntry, err error) error {
		if !dir.IsDir() {
			log.Printf("filename %s", dir.Name())
			log.Printf("%s", note)
			note = path
			info, err := dir.Info()
			if err != nil {
				return err
			}
			if maxModTime == 0 {
				maxModTime = info.ModTime().Unix()
				return nil
			}
			if maxModTime < info.ModTime().Unix() {
				maxModTime = info.ModTime().Unix()
				note = path
			}

		}

		return nil
	}

	err := filepath.WalkDir(noteDir, visit)
	log.Printf("file to edit: %s", note)
	if err != nil {
		fmt.Println(err)
	}

	// copy note

}

func Browse() {
	notesDir := fsutils.GetWorkingRepoDir()
	err := git.Browse(notesDir)
	if err != nil {
		return
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
