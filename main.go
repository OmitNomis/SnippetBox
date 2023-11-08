package main

import (
	"log"
	"net/http"
)

func home (w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/"{
		http.NotFound(w, r);
		return
	}
	w.Write([]byte ("Hello from SnippetBox"))
}
func showSnippet (w http.ResponseWriter, r *http.Request){
	w.Write([]byte ("Show all snippets..."))
}
func createSnippet (w http.ResponseWriter, r *http.Request){
	if (r.Method != "POST"){
		w.Header().Set("Allow", "POST")
		w.WriteHeader(405)
		w.Write([]byte ("Method not allowed "))
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