package main

import (
	"flag"
	"log"
	"net/http"
)

const (
	defaultListenAddr = ":8080"
)

func main() {
	var (
		listenAddr = flag.String("listen", defaultListenAddr, ":port to bind to")
	)
	flag.Parse()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("handling request", r)
	})
	log.Println("starting server at", *listenAddr)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
