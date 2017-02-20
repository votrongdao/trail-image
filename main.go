package trailimage

import (
	"html/template"
	"net/http"
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
}
