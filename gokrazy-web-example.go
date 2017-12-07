package main

import "log"
import "net/http"

func greeting(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Hello from a service run by gokrazy"))
}

func quitting(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Exiting the program!"))
}

func main() {
	http.HandleFunc("/greet", greeting)
	http.HandleFunc("/quit", quitting)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
