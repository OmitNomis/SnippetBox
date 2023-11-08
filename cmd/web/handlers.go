package main

import (
	"fmt"
	"net/http"
	"strconv"
)
 
func home(w http.ResponseWriter, r *http.Request){
	if (r.URL.Path != "/"){
		http.NotFound(w, r)
		return
	}
	w.Write([]byte ("This is home, hello World"))
}

func showSnippet (w http.ResponseWriter, r *http.Request){
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if (err != nil || id <1){
		http.NotFound(w, r)
		return
	}
	fmt.Fprintln(w, "ID for the snippet is :", id) 
}

func createSnippet (w http.ResponseWriter, r *http.Request){	
	if r.Method != "POST"	{
		http.Error(w, "Invalid Method called", 405)
		w.Header().Add("Allow", "POST")
		return
	}
	w.Write([]byte ("Create new Snippet"))
}