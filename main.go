package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"strings"
)

func archiveFiles(paths []string) {
	archive, err := os.Create("compressedFiles.zip")
	if(err != nil) {
		fmt.Println("Error creating zip archive", err)
	}
	defer archive.Close()

	myZip := zip.NewWriter(archive)

	for _, path := range paths {
		s := strings.Split(path, "/")
		file, err := os.Open(s[len(s)-1])
		if(err != nil) {
			fmt.Println("Error opening file", err)
		}
		defer file.Close()

		w, err := myZip.Create(path)
		if(err != nil) {
			fmt.Println("Error adding file to ZIP", err)
		}

		if _, err := io.Copy(w, file); err != nil {
			fmt.Println("Error copying file", err)
		}
	}
	myZip.Close()
}

func main() {
	paths := []string{"Example files 1\\Example file 1.txt",
					"Example files 1\\Example file 2.docx",
					"Example files 2\\Example file 3.xlsx"}
	
	archiveFiles(paths)
}