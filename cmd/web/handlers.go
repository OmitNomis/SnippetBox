package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
)
 
func home(w http.ResponseWriter, r *http.Request){
	if (r.URL.Path != "/"){
		http.NotFound(w, r)
		return
	}
	 
	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...) 
	if (err != nil) {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return;
	}
	err = ts.Execute(w, nil)
	
	if (err != nil){
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
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