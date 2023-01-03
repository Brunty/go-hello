package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hostname, err := os.Hostname()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/text")
		w.Write([]byte(hostname))
	})
	fmt.Printf("Starting server on :8180")
	if err := http.ListenAndServe(":8180", nil); err != nil {
		log.Fatal(err)
	}
}
