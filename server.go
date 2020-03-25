package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

var appVersion = "1.2" //Default/fallback version
var instanceNum int
var listenPort = "80" //Default port

func getFrontpage(w http.ResponseWriter, r *http.Request) {
	t := time.Now().UTC()
	fmt.Fprintf(w, "Hi, folks at London OpenHack! I'm instance %d running version %s of your application at %s\n", instanceNum, appVersion, t.Format("2006-01-02 15:04:05"))
}

func health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func getVersion(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s\n", appVersion)
}

func getEnvVars(w http.ResponseWriter, r *http.Request) {
	builder := strings.Builder{}
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		builder.WriteString(fmt.Sprintf("%s\n",pair[0]))
	}
	result := builder.String()
	fmt.Fprintf(w, "%s\n", result)
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	instanceNum = rand.Intn(1000)
	http.HandleFunc("/", getFrontpage)
	http.HandleFunc("/health", health)
	http.HandleFunc("/version", getVersion)
	http.HandleFunc("/env", getEnvVars)
	
	var listenAddress string
	if os.Getenv("LISTEN_PORT") != "" {
		listenAddress = fmt.Sprintf(":%s", os.Getenv("LISTEN_PORT"))
	} else {
		listenAddress = fmt.Sprintf(":%s", listenPort)
	}
	log.Fatal(http.ListenAndServe(listenAddress, nil))
}
