package main

import "net/http"

// See https://echo.labstack.com/cookbook/google-app-engine
func init() {
	admin := r.PathPrefix("/admin").Methods(http.MethodPost).Subrouter()
	r.HandleFunc("/admin/", test)
	admin.HandleFunc("/view/delete", test)
	admin.HandleFunc("/map/delete", test)
	admin.HandleFunc("/library/reload", test)
}
