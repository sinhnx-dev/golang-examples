package main

import (
	"compress/gzip"
	"io"
	"log"
	"os"
)

func main() {
	CompressFile("test.txt.gz")
	UncompressFile("test.txt.gz")
}
func CompressFile(gzipFile string) {
	// Create .gz file to write to
	outputFile, err := os.Create(gzipFile)
	if err != nil {
		log.Fatal(err)
	}

	// Create a gzip writer on top of file writer
	gzipWriter := gzip.NewWriter(outputFile)
	defer gzipWriter.Close()

	// When we write to the gzip writer
	// it will in turn compress the contents
	// and then write it to the underlying
	// file writer as well
	// We don't have to worry about how all
	// the compression works since we just
	// use it as a simple writer interface
	// that we send bytes to
	_, err = gzipWriter.Write([]byte("Gophers rule!\n"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Compressed data written to file.")
}
func UncompressFile(gzipFileName string) {
	// Open gzip file that we want to uncompress
	// The file is a reader, but we could use any
	// data source. It is common for web servers
	// to return gzipped contents to save bandwidth
	// and in that case the data is not in a file
	// on the file system but is in a memory buffer
	gzipFile, err := os.Open(gzipFileName)
	if err != nil {
		log.Fatal(err)
	}

	// Create a gzip reader on top of the file reader
	// Again, it could be any type reader though
	gzipReader, err := gzip.NewReader(gzipFile)
	if err != nil {
		log.Fatal(err)
	}
	defer gzipReader.Close()

	// Uncompress to a writer. We'll use a file writer
	outfileWriter, err := os.Create("unzipped.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer outfileWriter.Close()

	// Copy contents of gzipped file to output file
	_, err = io.Copy(outfileWriter, gzipReader)
	if err != nil {
		log.Fatal(err)
	}
}
