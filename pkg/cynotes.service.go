package cynotes

import (
	cryptoUtils "cynotes/pkg/crypto"
	"cynotes/pkg/editor"
	fsutils "cynotes/pkg/fs"
	"cynotes/pkg/git"
	promptUiUtils "cynotes/pkg/ui"
	"encoding/hex"
	"fmt"
	_ "github.com/samber/mo"
	"golang.org/x/exp/slices"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type MaxAttemptsError struct{}

func (m *MaxAttemptsError) Error() string {
	return "You have reached the maximum attempts, Bye!"
}

type Cynotes struct {
	fpcy cryptoUtils.FpCyUtil
}

func (cy Cynotes) renameTmpFile(tmpFilePath string, noteDir string) (string, error) {
	hash := cy.fpcy.GetMD5Hash(tmpFilePath).OrEmpty()
	hashStr := hex.EncodeToString(hash[:])
	secretNote := noteDir + "/" + hashStr
	err := os.Rename(tmpFilePath, secretNote)
	if err != nil {
		log.Printf("could not rename tmp file")
		return "", err
	}
	return secretNote, nil
}
func (cy Cynotes) push(notesDir string, secretNote string) {
	selection, _ := promptUiUtils.BasicPromptSelections("Do you want to push the changes?", []string{"Yes", "No"})
	if selection == "Yes" {
		git.Add(notesDir, secretNote)
		git.Commit(notesDir)
		git.Push(notesDir)
	}
}
func (cy Cynotes) encrypt(passphrase []byte, bytes []byte, secretNote string) error {
	ciphertext := cy.fpcy.EncryptAES(passphrase, bytes).RightOrEmpty()
	err := os.WriteFile(secretNote, ciphertext, 0755)
	if err != nil {
		log.Printf("Failed reading data from file: %s", err)
		return err
	}
	return nil
}
func (cy Cynotes) decrypt(attempts uint8, cyBytes []byte) ([]byte, string, error) {
	var _bytes []byte
	var passphrase string
	if attempts > 0 {
		passphrase, _ = promptUiUtils.PromptPasswdInput()
		bytes, isRight := cy.fpcy.DecryptAES([]byte(passphrase), cyBytes).Right()
		if !isRight {
			attempts--
			log.Printf("Wrong passphrase, try again")
			return cy.decrypt(attempts, cyBytes)
		}
		_bytes = bytes
		return _bytes, passphrase, nil
	}
	return nil, "", &MaxAttemptsError{}
}

// External API

func InitCYNotes(user string, repo string) error {
	// Greet
	uname, err := fsutils.GetUserName()
	if err != nil {
		return err
	}
	fmt.Printf("Hey %s\n", uname)

	path, err := fsutils.NormalizeCYNotesPath(uname)
	localRepoPath := path + "/" + repo
	fmt.Printf("local repo path: %s \n", localRepoPath)
	if !fsutils.IsPathExists(localRepoPath) {
		// Create docs base folder
		fmt.Printf("Creating directory: %s \n", path)
		err = os.Mkdir(path, 0755)
		if err != nil {
			fmt.Printf("Error: %s \n", err)
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

func List() []string {
	notes := []string{}
	visit := func(path string, dir fs.DirEntry, err error) error {
		if !strings.Contains(path, ".git") && !slices.Contains(notes, path) {
			notes = append(notes, "\n"+path)
		}
		return nil
	}

	path, _ := fsutils.GetCYNotesPath()
	err := filepath.WalkDir(path, visit)
	if err != nil {
		fmt.Println(err)
	}
	return notes
}

func New(name string) error {
	notesDir := fsutils.GetWorkingRepoDir()
	cy := Cynotes{}
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
	secretNote, _ := cy.renameTmpFile(tmpFilePath, notesDir+"/"+name)

	// Encrypt
	log.Printf("Enter passphrase for encrypting note")
	passphrase, _ := promptUiUtils.PromptPasswdInput()

	bytes, err := os.ReadFile(secretNote)
	if err != nil {
		log.Printf("Failed reading data from file: %s", err)
		return err
	}

	_ = cy.encrypt([]byte(passphrase), bytes, secretNote)
	cy.push(notesDir, secretNote)
	return nil
}

func Edit(name string) {
	cy := Cynotes{}
	notesDir := fsutils.GetWorkingRepoDir()
	noteDir := notesDir + "/" + name
	log.Printf(noteDir)

	_note := ""
	var maxModTime int64 = 0

	visit := func(path string, dir fs.DirEntry, err error) error {
		if !dir.IsDir() {
			info, err := dir.Info()
			if err != nil {
				return err
			}

			if maxModTime < info.ModTime().Unix() { // grab the latest update
				maxModTime = info.ModTime().Unix()
				_note = path
			}

		}

		return nil
	}

	err := filepath.WalkDir(noteDir, visit)
	log.Printf("note to edit: %s", _note)
	if err != nil {
		fmt.Println(err)
	}

	// copy note
	cyBytes, err := os.ReadFile(_note)
	if err != nil {
		fmt.Printf("Failed reading data from file: %s", err)
	}

	if len(cyBytes) == 0 {
		log.Printf("The file seems to be empty")
		os.Exit(1)
	}
	// try to decrypt encrypted note by user passwd attempts
	bytes, passphrase, err := cy.decrypt(3, cyBytes)

	if err != nil {
		log.Printf("Error: %s", err)
		os.Exit(1)
	}

	tmpFile := noteDir + "/tmp_file"
	err = os.WriteFile(tmpFile, bytes, 0755)
	if err != nil {
		log.Printf("Failed reading data from file: %s", err)
	}

	_ = editor.ViEdit(tmpFile)
	secretNote, _ := cy.renameTmpFile(tmpFile, notesDir+"/"+name)

	editedBytes, err := os.ReadFile(secretNote)
	_ = cy.encrypt([]byte(passphrase), editedBytes, secretNote)
	cy.push(notesDir+"/"+name, secretNote)

}

func Browse() {
	notesDir := fsutils.GetWorkingRepoDir()
	err := git.Browse(notesDir)
	if err != nil {
		log.Printf("Failed browsing to repo: %s", err)
		return
	}
}

func Read(note string) {
	var revs []string
	visit := func(path string, dir fs.DirEntry, err error) error {
		if !dir.IsDir() {
			revs = append(revs, path)
		}
		return nil
	}

	path := fsutils.GetWorkingRepoDir()
	err := filepath.WalkDir(path+"/"+note, visit)
	if err != nil {
		return
	}
	selection, _ := promptUiUtils.BasicPromptSelections("Select a revision to read", revs)
	cyBytes, err := os.ReadFile(selection)
	bytes, _, err := Cynotes{}.decrypt(3, cyBytes)
	log.Print(string(bytes[:]))
}
