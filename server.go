package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var appVersion = "1.2" //Default/fallback version
var instanceNum int

// Old catch all handler
// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:]) //substring after '/' character
// }

// func getenv(key, fallback string) string {
// 	value := os.Getenv(key)
// 	if len(value) == 0 {
// 		return fallback
// 	}
// 	return value
// }

func getFrontpage(w http.ResponseWriter, r *http.Request) {
	t := time.Now().UTC()
	fmt.Fprintf(w, "Hi! I'm CSE folks! I'm instance %d running version %s of your application at %s\n", instanceNum, appVersion, t.Format("2006-01-02 15:04:05"))
}

func health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func getVersion(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s\n", appVersion)
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	instanceNum = rand.Intn(1000)
	http.HandleFunc("/", getFrontpage)
	http.HandleFunc("/health", health)
	http.HandleFunc("/version", getVersion)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
