// mkdir 2-chi-hello-world && cd 2-chi-hello-world
// go mod init helloapi-chi
// go mod tidy ("to add module requirements and sums")
// go run .
// curl localhost:8090
// curl localhost:8090/hello

package main

import (
	"fmt"
	"log"
	"net/http"

  "github.com/go-chi/chi/v5"
  "github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", handleIndex)
	r.Get("/hello", handleHello)
	
	fmt.Println("Listening at :8090...")
	
	// Note that handler = r
	err := http.ListenAndServe(":8090", r)
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
