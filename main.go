package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	target := "http://localhost:1106"
	// target := "http://localhost:8283"
	targetURL, err := url.Parse(target)
	if err != nil {
		log.Fatalf("Failed to parse target URL: %v", err) // dadata "6aa6c4610172cd3ac53226f64516c1b6969ed7df"
	}

	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	})

	err = http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatalf("Failed to start server on port 80: %v", err)
	}
}
