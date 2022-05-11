package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/onlyPost", handlePost)

	http.ListenAndServe(":8080", nil)
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	log.Println("This is ", r.FormValue("fruit"))
	w.Write([]byte("This is " + r.FormValue("fruit")))
}
