package fsutils

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"strings"
)

/**
  Currently only supporting macOS

	Todo:
		replace string concatenation with path lib
		choose either fmt or log instead of both
*/

func getUserName() (string, error) {
	_user, err := user.Current()
	return _user.Username, err
}

// Assuming there's only one person/user which uses the computer (pc)
// Following protects a case where the profile/userpath name isn't
// perfectly matching the os username e.g. os_user == liron but userpath == liron_1
func getCYNotesPath(uname string) (string, error) {
	files, err := os.ReadDir("/Users")
	var name string
	for _, file := range files {
		if strings.HasPrefix(file.Name(), uname) {
			name = file.Name()
			break
		}
	}
	return "/Users/" + name + "/.cynotes", err
}

func isPathExists(path string) bool {
	if _, err := os.Stat(path); err == nil || os.IsExist(err) {
		fmt.Printf(" Base docs folder: %s\n", path)
		return true
	}
	return false
}

func InitCYNotes() error {
	// Greet
	uname, err := getUserName()
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("Hey %s\n", uname)

	path, err := getCYNotesPath(uname)
	if !isPathExists(path) {
		// Create docs base folder
		err = os.Mkdir(path, 0755)
		if err != nil {
			log.Fatalf(err.Error())
		}
	}
	return err
}

func GetCYNotesPath() (string, error) {
	uname, err := getUserName()
	path, err := getCYNotesPath(uname)
	return path, err
}

func CreateFolder(absPath string) {
	if !isPathExists(absPath) {
		fmt.Printf(" Creating %s \n", absPath)
		err := os.Mkdir(absPath, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func CreateNoteFolder(filename string) string {
	path, _ := GetCYNotesPath()
	var fPath = path + "/" + filename
	CreateFolder(fPath)
	return fPath
}

func ModTimeUnix(filename string) (int64, error) {
	file, err := os.Stat(filename)
	return file.ModTime().Unix(), err
}

func ExtractFilename(filepath string) string {
	slice := strings.Split(filepath, "/")
	return slice[len(slice)-1]
}
