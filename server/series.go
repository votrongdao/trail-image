package main

import "net/http"

func init() {
	series := r.PathPrefix("/" + TOKEN_SERIES).Methods(http.MethodGet).Subrouter()
	series.HandleFunc("/", test)
	series.HandleFunc("/map", test)
	series.HandleFunc("/map/"+TOKEN_PHOTO_ID, test)

}
