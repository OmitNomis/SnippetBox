package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)


func main (){
	// Get port from cli
	addr := flag.String( "addr", ":4000", "HTTP network address")
	flag.Parse()

	// create custom loggers
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ldate|log.Lshortfile)


	mux:= http.NewServeMux();
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	// creating fileserver for uploading all static files
	fileserver := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileserver))

	svr := &http.Server{
		Addr: *addr,
		ErrorLog: errorLog,
		Handler: mux,
	}

	infoLog.Println("Starting server on port", *addr)

	err := svr.ListenAndServe()
	errorLog.Fatal(err)
	
}