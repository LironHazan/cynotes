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

func renameTmpFile(tmpFilePath string, noteDir string) (string, error) {
	hash := cryptoUtils.GetMD5Hash(tmpFilePath)
	hashStr := hex.EncodeToString(hash[:])
	secretNote := noteDir + "/" + hashStr
	err := os.Rename(tmpFilePath, secretNote)
	if err != nil {
		log.Printf("could not rename tmp file")
		return "", err
	}
	return secretNote, nil
}
func push(notesDir string, secretNote string) {
	selection, _ := promptUiUtils.BasicPromptSelections("Do you want to push the changes?", []string{"Yes", "No"})
	if selection == "Yes" {
		git.Add(notesDir, secretNote)
		git.Commit(notesDir)
		git.Push(notesDir)
	}
}
func encrypt(passphrase []byte, bytes []byte, secretNote string) error {
	ciphertext, err := cryptoUtils.EncryptAES(passphrase, bytes)
	err = os.WriteFile(secretNote, ciphertext, 0755)
	if err != nil {
		log.Printf("Failed reading data from file: %s", err)
		return err
	}
	return nil
}

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
	secretNote, _ := renameTmpFile(tmpFilePath, notesDir+"/"+name)

	// Encrypt
	log.Printf("Enter passphrase for encrypting note")
	passphrase, _ := promptUiUtils.PromptPasswdInput()

	bytes, err := os.ReadFile(secretNote)
	if err != nil {
		log.Printf("Failed reading data from file: %s", err)
		return err
	}

	_ = encrypt([]byte(passphrase), bytes, secretNote)
	push(notesDir, secretNote)
	return nil
}

func EditNote(name string) {
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
	bytes, err := os.ReadFile(note)
	if err != nil {
		fmt.Printf("Failed reading data from file: %s", err)
	}

	passphrase, _ := promptUiUtils.PromptPasswdInput()
	text, err := cryptoUtils.DecryptAES([]byte(passphrase), bytes)

	tmpFile := noteDir + "/tmp_file"
	err = os.WriteFile(tmpFile, text, 0755)
	if err != nil {
		log.Printf("Failed reading data from file: %s", err)
	}

	_ = editor.ViEdit(tmpFile)
	secretNote, _ := renameTmpFile(tmpFile, notesDir+"/"+name)
	_ = encrypt([]byte(passphrase), bytes, secretNote)
	push(notesDir+"/"+name, secretNote)

}

func Browse() {
	notesDir := fsutils.GetWorkingRepoDir()
	err := git.Browse(notesDir)
	if err != nil {
		return
	}
}
