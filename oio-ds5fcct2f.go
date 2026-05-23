package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

type server struct{}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}

func main() {
	srv := &server{}
	log.Fatal(http.ListenAndServe(":8080", srv))
}