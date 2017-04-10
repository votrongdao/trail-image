package main

import "net/http"

func init() {
	tag := r.PathPrefix("/photo-tag").Methods(http.MethodGet).Subrouter()
	tag.HandleFunc("/", test)
	tag.HandleFunc("/{"+PATH_PHOTO_TAG+"}", test)
	tag.HandleFunc("/search/{"+PATH_PHOTO_TAG+"}", test)
}
