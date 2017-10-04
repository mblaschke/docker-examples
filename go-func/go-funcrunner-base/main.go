package main

import (
	"fmt"
	"os"
	"log"
	"net/http"
)

var (
	logger *log.Logger
	mux *http.ServeMux
)

func main() {
	logger = log.New(os.Stdout, "logger: ", log.Lshortfile)

	fmt.Println("Starting go http server...")
	mux = http.NewServeMux()
	mux.HandleFunc("/", apifunction)
	http.ListenAndServe(":8000", mux)
}
