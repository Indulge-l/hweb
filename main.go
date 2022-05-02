package main

import "net/http"

func main() {
	fs := http.FileServer(http.Dir("./"))
	http.Handle("/static", http.StripPrefix("/static", fs))
	panic(http.ListenAndServe(":8080", nil))
}
