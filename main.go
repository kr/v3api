package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

func main() {
	log.SetFlags(0)
	if len(os.Args) != 2 {
		log.Fatal("Usage: v3api addr")
	}
	addr := os.Args[1]
	http.Handle("/", &httputil.ReverseProxy{Director: Director})
	log.Fatal(http.ListenAndServe(addr, nil))
}

func Director(r *http.Request) {
	r.Host = "api.heroku.com"
	r.URL.Scheme = "https"
	r.URL.Host = "api.heroku.com"
	r.Header.Set("Accept", "application/vnd.heroku+json; version=3")
}
