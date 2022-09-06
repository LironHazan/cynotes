package fsutils

import (
	"encoding/json"
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

type InitData struct {
	RepoName string
}

func CreateInitFile(repo string) {
	path, _ := GetCYNotesPath()
	data := InitData{RepoName: repo}
	file, _ := json.MarshalIndent(data, "", " ")
	_ = os.WriteFile(path+"/.init.json", file, 0644)
}

func GetRepoName() string {
	path, _ := GetCYNotesPath()
	name := path + "/.init.json"
	file, _ := os.ReadFile(name)

	data := InitData{}

	_ = json.Unmarshal([]byte(file), &data)
	fmt.Println("RepoName: ", data.RepoName)
	return data.RepoName
}

func GetWorkingRepoDir() string {
	repo := GetRepoName()
	path, _ := GetCYNotesPath()
	return path + "/" + repo
}

func GetUserName() (string, error) {
	_user, err := user.Current()
	return _user.Username, err
}

// Assuming there's only one person/user which uses the computer (pc)
// Following protects a case where the profile/userpath name isn't
// perfectly matching the os username e.g. os_user == liron but userpath == liron_1
func NormalizeCYNotesPath(uname string) (string, error) {
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

func IsPathExists(path string) bool {
	if _, err := os.Stat(path); err == nil || os.IsExist(err) {
		return true
	}
	return false
}

func GetCYNotesPath() (string, error) {
	uname, err := GetUserName()
	path, err := NormalizeCYNotesPath(uname)
	return path, err
}

func CreateFolder(absPath string) {
	if !IsPathExists(absPath) {
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
