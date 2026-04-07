package main

import (
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func helloHandler1(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func helloHandler2(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func helloHandler3(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func helloHandler4(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/api/hello", helloHandler)
	http.ListenAndServe(":8080", nil)

	http.HandleFunc("/api/hello1", helloHandler1)
	http.ListenAndServe(":8080", nil)

	http.HandleFunc("/api/hello2", helloHandler2)
	http.ListenAndServe(":8080", nil)

	http.HandleFunc("/api/hello3", helloHandler3)
	http.ListenAndServe(":8080", nil)

	http.HandleFunc("/api/hello4", helloHandler4)
	http.ListenAndServe(":8080", nil)
}
