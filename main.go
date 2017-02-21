package main

import (
	"html/template"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

func init() {
	http.HandleFunc("/", hello)
}

// see https://hackernoon.com/golang-template-1-bcb690165663
func hello(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/post.html")
	if err != nil {
		w.Write([]byte("Unable to find template"))
	} else {
		t.Execute(w, "Hello World!")
	}

	ctx := appengine.NewContext(r)
	q := datastore.NewQuery("Greeting").Ancestor(guestbookKey(ctx)).Order("-Date").Limit(10)
}
