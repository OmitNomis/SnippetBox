package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)
 
func (app *application) home(w http.ResponseWriter, r *http.Request){
	if (r.URL.Path != "/"){
		app.notFound(w)
		return
	}
	 
	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...) 
	if (err != nil) {
		app.serverError(w, err)
		
		return;
	}
	err = ts.Execute(w, nil)
	
	if (err != nil){
		app.serverError(w, err)
	}
}

func (app *application) showSnippet (w http.ResponseWriter, r *http.Request){
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if (err != nil || id <1){
		app.notFound(w)
		return
	}
	fmt.Fprintln(w, "ID for the snippet is :", id) 
}

func (app *application) createSnippet (w http.ResponseWriter, r *http.Request){	
	if r.Method != "POST"	{
		http.Error(w, "Invalid Method called", 405)
		w.Header().Add("Allow", "POST")
		return
	}
	w.Write([]byte ("Create new Snippet"))
}