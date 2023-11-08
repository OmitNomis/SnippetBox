package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home (w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/"{
		http.NotFound(w, r);
		return
	}
	w.Write([]byte ("Hello from SnippetBox"))
}
func showSnippet (w http.ResponseWriter, r *http.Request){
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if (err != nil || id < 1){
		http.NotFound(w, r)
		return;
	}
	fmt.Println("id from queryParam", id)
	w.Write([]byte ("Show Snippet info..."))
}
func createSnippet (w http.ResponseWriter, r *http.Request){
	if (r.Method != "POST"){
		w.Header().Set("Allow", "POST")
		http.Error(w, "Method not allowed", 405)
		return;
	}
	w.Write([]byte ("Create a new snippet"))
}

func main () {
	
	mux:= http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)
	log.Println("Starting Server from :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}