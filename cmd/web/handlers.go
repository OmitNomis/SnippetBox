package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)
 
func (app *application) home(w http.ResponseWriter, r *http.Request){
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
		app.errorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return;
	}
	err = ts.Execute(w, nil)
	
	if (err != nil){
		app.errorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func (app *application) showSnippet (w http.ResponseWriter, r *http.Request){
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if (err != nil || id <1){
		http.NotFound(w, r)
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