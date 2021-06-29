package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Minha primeira app Golang!")
	})
	http.ListenAndServe(":1337", nil)
}
