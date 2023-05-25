// package main

// import (
// 	"net/http"
// )

// func main() {
// 	mux := http.NewServeMux()
// 	fs := http.FileServer(http.Dir("public"))
// 	mux.Handle("/", fs)
// 	http.ListenAndServe(":8080", mux)
// }

package main

import (
	"fmt"
	"log"
	"net/http"
)

type messageHandler struct {
	message string
}

func (m *messageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, m.message)
}
func main() {
	mux := http.NewServeMux()
	mh1 := &messageHandler{"Welcome to Go Web Development"}
	mux.Handle("/welcome", mh1)
	mh2 := &messageHandler{"<h1>net/http is awesome</h1>"}
	mux.Handle("/message", mh2)
	log.Println("Listening...")
	http.ListenAndServe(":8080", mux)
}
