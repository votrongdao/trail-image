package main

import "net/http"

func init() {
	c := r.PathPrefix("/admin").Methods(http.MethodGet).Subrouter()
	c.HandleFunc("/", test)
	c.HandleFunc("/{"+PATH_CATEGORY+"}", test)
}
