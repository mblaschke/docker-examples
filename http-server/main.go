package main

import (
	"fmt"
	"time"
	"io"
	"os"
	"log"
	"net"
	"net/http"
)

var (
	logger *log.Logger 
)

func hello(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	currentDate := time.Now().Format(time.RFC850)
	io.WriteString(w, fmt.Sprintf("Hello world! I'm %s and it's %s", hostname, currentDate))

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err == nil {
		userIP := net.ParseIP(ip)
		logger.Println(fmt.Sprintf("%s: Request from %v to %v", currentDate, userIP, hostname))
	}
}

func main() {
	logger = log.New(os.Stdout, "logger: ", log.Lshortfile)

	fmt.Println("Starting go http server...")
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	http.ListenAndServe(":80", mux)
}

