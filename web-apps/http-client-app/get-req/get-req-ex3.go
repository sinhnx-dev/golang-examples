package main

import (
	"log"
	"net/http"
	"os"
)

func main() {

	r, err := http.Get("http://sinhnx.dev")

	if err != nil {
		log.Fatal(err)
	}

	defer r.Body.Close()

	f, err := os.Create("index.html")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err = f.ReadFrom(r.Body)

	if err != nil {
		log.Fatal(err)
	}
}
