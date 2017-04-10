package main

import "net/http"

func init() {
	series := r.PathPrefix("/admin").Methods(http.MethodGet).Subrouter()
	series.HandleFunc("/", test)
	series.HandleFunc("/map", test)
	series.HandleFunc("/map/{"+RE_PHOTO_ID+"}", test)

}
