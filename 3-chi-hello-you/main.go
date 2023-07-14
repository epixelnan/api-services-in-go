// mkdir 3-chi-hello-you && cd 3-chi-hello-you
// go mod init helloapi-chi-param
// go mod tidy ("to add module requirements and sums")
// go run .
// curl localhost:8091
// curl localhost:8091/hello/Nandakumar

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
	r.Get("/hello/{name}", handleHello)
	
	fmt.Println("Listening at :8091...")
	
	// Note that handler = r
	err := http.ListenAndServe(":8091", r)
	if err != nil {
		log.Fatal(err)
	}

	// Will never reach
	fmt.Println("Exiting.")
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, " + chi.URLParam(r, "name") + "!\n"))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome!\n"))
}
