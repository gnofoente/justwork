package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.String()
		log.Printf("url: %s\tmethod: %s\n", url, r.Method)
		resp, err := http.Get(url)
		if err != nil {
			log.Printf("failed fetching from %s: %s", url, err)
		}
		defer resp.Body.Close()
		io.Copy(w, resp.Body)
	})

	http.ListenAndServe("127.0.0.1:8080", nil)
}
