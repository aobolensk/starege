package main

import (
	"flag"
	"starege/server"
)

type Args struct {
	port     int
	rootUrl  string
	serveDir string
}

func parseArgs() Args {
	args := new(Args)
	flag.IntVar(&args.port, "port", 8080, "port number")
	flag.StringVar(&args.rootUrl, "rootUrl", "starege", "URL to serve main page at")
	flag.StringVar(&args.serveDir, "serveDir", "", "Directory to serve")
	flag.Parse()
	return *args
}

func main() {
	args := parseArgs()
	server.Serve("", args.port, args.rootUrl, args.serveDir)
}
