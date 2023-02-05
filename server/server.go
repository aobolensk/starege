package server

import (
	"fmt"
	"log"
	"net/http"
)

func Serve(ip string, port int, rootUrl string, serveDir string) {
	if serveDir != "" {
		files := http.StripPrefix("/"+rootUrl+"/data/", http.FileServer(http.Dir(serveDir)))
		http.Handle("/"+rootUrl+"/data/", files)
	}
	ipPort := fmt.Sprintf(":%d", port)
	log.Printf("Listening and serving starege on %s", ipPort)
	log.Fatal(http.ListenAndServe(ipPort, nil))
}
