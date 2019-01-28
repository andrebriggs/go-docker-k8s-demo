package main

import (
	"fmt"
	"log"
	"net/http"
)

const version string = "1.2"

// Old catch all handler
// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:]) //substring after '/' character
// }

func getFrontpage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Congratulations! Version %s of your application is running on Kubernetes.", version)
}

func health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func getVersion(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s\n", version)
}

func main() {
	http.HandleFunc("/", getFrontpage)
	http.HandleFunc("/health", health)
	http.HandleFunc("/version", getVersion)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
