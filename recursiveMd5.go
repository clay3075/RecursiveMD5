package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	searchDir := os.Args[1]
	if searchDir == "" {
		searchDir = "."
	}

	fileList := []string{}
	filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		if !f.IsDir() {
			fileList = append(fileList, path)
		}

		return nil
	})
	var fileContents strings.Builder
	for _, file := range fileList {
		fmt.Fprintf(&fileContents, "\"%s\",\"%x\"\n", strings.Replace(file, searchDir+"\\", "", 1), getMD5(file))
	}

	writeToMD5File(searchDir, fileContents.String())
}

func writeToMD5File(dir string, contents string) {
	f, e := os.Create(dir + "\\md5.txt")
	check(e)
	defer f.Close()
	_, e = f.WriteString(contents)
	check(e)
}

func getMD5(file string) []byte {
	f, err := os.Open(file)
	check(err)
	defer f.Close()

	h := md5.New()
	_, err = io.Copy(h, f)
	check(err)

	return h.Sum(nil)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
