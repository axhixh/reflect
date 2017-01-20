package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		header := w.Header()
		header.Set("Content-Type", "text/plain")

		fmt.Fprintln(w, "Headers")
		for key, value := range r.Header {
			fmt.Fprintln(w, key+": ", value)
		}
		fmt.Fprintln(w, "Method: ", r.Method)
		fmt.Fprintln(w, "URL: ", r.URL)
		fmt.Fprintln(w, "Host: ", r.Host)
		fmt.Fprintln(w, "Remote Address: ", r.RemoteAddr)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Starting Reflection at %s", port)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("Error starting Reflection. ", err)
	}
}
