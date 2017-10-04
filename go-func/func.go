package main

import (
	"io"
	"net/http"
)

func apifunction(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "Hello world! This is my api function")
    logger.Println("Called my api function")
}
