package blog

import (
	"html/template"
	"net/http"
)

func init() {
	http.HandleFunc("/", test)
}

// see https://hackernoon.com/golang-template-1-bcb690165663
func test(w http.ResponseWriter, r *http.Request) {
	// c := appengine.NewContext(r)
	t, err := template.ParseFiles("templates/post.html")
	if err != nil {
		w.Write([]byte("Unable to find template"))
	} else {
		t.Execute(w, "Hello World!")
	}

	//ctx := appengine.NewContext(r)
	// q := datastore.NewQuery("Greeting").Ancestor(guestbookKey(ctx)).Order("-Date").Limit(10)
}
