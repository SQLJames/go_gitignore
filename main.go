package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-git/go-git"
)

func Warning(format string, args ...interface{}) {
	fmt.Printf("\x1b[36;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}

func CheckArgs(arg ...string) {
	if len(os.Args) < len(arg)+1 {
		Warning("Usage: %s %s", os.Args[0], strings.Join(arg, " "))
		os.Exit(1)
	}
}

func CheckIfError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func GetFiles(directory string) ([]string, error) {
	var files []string
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}
func getExe() (string, error) {
	executable, err := os.Executable()
	executablelist := strings.Split(executable, "\\")
	lenminusone := len(executablelist) - 1
	return strings.Join(executablelist[0:lenminusone], "\\"), err
}
func main() {
	CheckArgs("<LanguageType>")
	LanguageType := os.Args[1]
	LanguageType = strings.ToLower(LanguageType)
	Exepath, err := getExe()
	fmt.Println(Exepath)
	//TODO: implement this in a tempdir location
	directory := "gitignore"
	os.Mkdir(directory, 0777)
	defer os.RemoveAll(directory)
	_, err = git.PlainClone(directory, false, &git.CloneOptions{
		URL:      "https://github.com/github/gitignore",
		Progress: os.Stdout,
	})
	CheckIfError(err)
	files, err := GetFiles(filepath.Join(Exepath, directory))
	CheckIfError(err)
	os.Mkdir("results", 0777)
	for _, file := range files {
		file = strings.ToLower(file)
		fileinfo, err := os.Stat(file)
		CheckIfError(err)
		if strings.Contains(file, LanguageType) {
			filename := fileinfo.Name()
			Exportlocation := filepath.Join(Exepath, "results", filename)

			fmt.Printf("%s moving to %s", file, Exportlocation)
			os.Rename(file, Exportlocation)
		}
	}
}
