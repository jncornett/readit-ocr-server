package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

const (
	defaultListenAddr = ":8080"
)

func main() {
	var (
		listenAddr = flag.String("listen", envDefault("LISTEN", defaultListenAddr), ":port to bind to")
	)
	flag.Parse()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("handling request", r)
	})
	log.Println("starting server at", *listenAddr)
	log.Fatal(http.ListenAndServe(*listenAddr, nil))
}

func envDefault(key, dflt string) string {
	value := os.Getenv(key)
	if value == "" {
		return dflt
	}
	return value
}
