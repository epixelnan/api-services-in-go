// mkdir 1-hello-world && cd 1-hello-world
// go mod init helloapi
// go mod tidy ("to add module requirements and sums")
// go run .
// curl localhost:8080
// curl localhost:8080/hello

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Adds to DefaultServeMux
	// Note: you can use anonymous functions
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/hello", handleHello)
	
	fmt.Println("Listening at :8080...")
	
	// Address and handler; handler = nil means DefaultServeMux
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

	// Will never reach
	fmt.Println("Exiting.")
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("world!\n"))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome!\n"))
}
