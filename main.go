package main

import (
	"fmt"
	"net/http"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(404)
		w.Write([]byte("page not found!"))
		return
	}
	if r.Method == "GET" {
		fmt.Fprint(w, "welcome to the index of this page.")
	} else {
		w.WriteHeader(405)
		w.Write([]byte("Method not allowed"))
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleIndex)
	fmt.Println("started listening")
	err := http.ListenAndServe(":5060", mux)
	if err != nil {
		fmt.Printf("server failed %s", err.Error())
	}
}
