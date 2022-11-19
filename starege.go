package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

type Args struct {
	port     int
	serveDir string
}

func parseArgs() Args {
	args := new(Args)
	flag.IntVar(&args.port, "port", 8080, "port number")
	flag.StringVar(&args.serveDir, "serveDir", "", "Directory to serve")
	flag.Parse()
	return *args
}

func main() {
	args := parseArgs()
	if args.serveDir != "" {
		files := http.StripPrefix("/data/", http.FileServer(http.Dir(args.serveDir)))
		http.Handle("/data/", files)
	}
	ipPort := fmt.Sprintf(":%d", args.port)
	log.Printf("Listening and serving starege on %s", ipPort)
	log.Fatal(http.ListenAndServe(ipPort, nil))
}
