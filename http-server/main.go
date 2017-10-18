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

const (
	VERSION = "1.0"
)

func hello(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	currentDate := time.Now().Format(time.RFC850)
	io.WriteString(w, fmt.Sprintf("Hello world! I'm %s in version %s and it's %s", hostname, VERSION, currentDate))

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err == nil {
		userIP := net.ParseIP(ip)
		logger.Println(fmt.Sprintf("%s: Request from %v to %v", currentDate, userIP, hostname))
	}
}

func main() {
	logger = log.New(os.Stdout, "logger: ", log.Lshortfile)

	fmt.Println(fmt.Sprintf("Starting go http server version %s...", VERSION))
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	http.ListenAndServe(":80", mux)
}

