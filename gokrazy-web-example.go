package main

import "log"
import "net/http"

func Greeting(w http.ResponseWriter, req *http.Request) {
    w.Header().Set("Content-Type", "text/plain")
    w.Write([]byte("Hello from a service run by gokrazy"))
}

func main() {
    http.HandleFunc("/greet", Greeting)
    log.Fatal(http.ListenAndServe(":8080",nil))
}
