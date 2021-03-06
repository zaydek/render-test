package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s %s\n", r.Method, r.URL.String())
		fmt.Fprintln(w, "Hello, world!")
	})
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}
