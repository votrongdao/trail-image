package main

import "net/http"

func init() {
	post := r.PathPrefix("/admin").Methods(http.MethodGet).Subrouter()
	post.HandleFunc("/", test)
	post.HandleFunc("/map", test)
	post.HandleFunc("/gpx", test)
	post.HandleFunc("/map/{"+RE_PHOTO_ID+"}", test)
	post.HandleFunc("/geo.json", test)
}
