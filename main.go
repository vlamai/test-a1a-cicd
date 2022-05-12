package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello the current time is %s", time.Now())
	})
	log.Println("Start server on http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}
