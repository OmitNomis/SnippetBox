package main

import (
	"flag"
	"log"
	"net/http"
)


func main (){

	addr := flag.String( "addr", ":4000", "HTTP network address")
	flag.Parse()
	mux:= http.NewServeMux();
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	fileserver := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileserver))
	log.Println("Srarting server on port", *addr);
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
	
}