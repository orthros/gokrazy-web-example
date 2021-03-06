package main

import "log"
import "net/http"
import "os"

func greeting(w http.ResponseWriter, req *http.Request) {
	log.Print("Greeting!")
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Hello from a service run by gokrazy"))
}

func quitting(w http.ResponseWriter, req *http.Request) {
	log.Print("Quitting!")
	os.Exit(0)
}

func main() {
	log.Print("Coming online!")
	http.HandleFunc("/greet", greeting)
	http.HandleFunc("/quit", quitting)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
