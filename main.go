package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal("Could not get hostname: ", err)
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {

		if req.URL.Path != "/" {
			http.NotFound(w, req)
			return
		}
		fmt.Fprintf(w, "Hostname: %s", hostname)
	})

	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("ListanAndServe: ", err)
	}
}
