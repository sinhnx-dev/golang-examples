package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Open a new file for writing only
	file, err := os.OpenFile(
		"test.txt",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Write bytes to file
	byteSlice := []byte("Bytes!\n")
	bytesWritten, err := file.Write(byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Wrote %d bytes.\n", bytesWritten)

	file, err = os.OpenFile(
		"test.txt",
		os.O_RDONLY,
		0666,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read bytes to file
	bytesNo, err := file.Read(byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	if bytesNo > 0 {
		fmt.Printf("Read from file: \"%s\"", string(byteSlice))
	}
}
