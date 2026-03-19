package mobileappreactnative

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func isFileExists(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}
	return true
}

func isDirExists(filePath string) bool {
	s, err := os.Stat(filePath)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func IsEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

func getExecutablePath() string {
	//Get the path to the current executable
	exe, err := os.Executable()
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return filepath.Dir(exe)
}

func ReadFileToString(filePath string) (string, error) {
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(b), nil
}