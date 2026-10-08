package main

import "net/http"

func main() {
	println("Serving at http://localhost:8080 (Ctrl+C to stop)")
	if err := http.ListenAndServe(":8080", http.FileServer(http.Dir("public"))); err != nil {
		panic(err)
	}
}
