package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, Welcome to AzAPIBackendAuthz!")
}

func signInHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Sign In Successful")
}

func signOutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Sign Out Successful")
}

func dataHandlerInsensitive(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "API data response post auth - Insensitive")
}

func dataHandlerSensitive(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "API data response post auth - Sensitive")
}

func main() {
	hostIp := "0.0.0.0"
	hostPort := 6666
	completeHost := hostIp + ":" + strconv.Itoa(hostPort)
	fmt.Println("Starting server on " + completeHost)
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/signin", signInHandler)
	http.HandleFunc("/signout", signOutHandler)
	http.HandleFunc("/api/insensitive", dataHandlerInsensitive)
	http.HandleFunc("/api/sensitive", dataHandlerSensitive)
	log.Fatal(http.ListenAndServe(completeHost, nil))
}
