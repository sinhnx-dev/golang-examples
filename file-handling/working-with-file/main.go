package main

import (
	"log"
	"os"
)

var (
	newFile *os.File
	err     error
)

func main() {
	// Create Empty File
	newFile, err = os.Create("./text-file.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(newFile)
	newFile.Close()

	//Rename and Move a File
	originalPath := "text-file.txt"
	newPath := "test-file.txt"
	err = os.Rename(originalPath, newPath)
	if err != nil {
		log.Fatal(err)
	}

	//Delete a file
	err = os.Remove("test-file.txt")
	if err != nil {
		log.Fatal(err)
	}
}
