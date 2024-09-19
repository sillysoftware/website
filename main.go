// main.go defines the main logic for the Silly Software Foundation website
//
//	Copyright (C) 2024-2024 Silly Software Foundation
//
//	This file is part of the Silly Software Foundation
//
//	This website is free software; you can redistribute it and/or midify it under the BSD 3-Clause Licence.
package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		err := tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
	http.HandleFunc("/", h1)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
