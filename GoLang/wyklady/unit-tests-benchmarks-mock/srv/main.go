package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "World")
	})

	if err := http.ListenAndServe(":9999", nil); err != nil {
		log.Fatal(err)
	}
}
