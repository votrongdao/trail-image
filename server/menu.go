package main

import (
	"html/template"
	"net/http"

	"trailimage.com/format"

	"github.com/labstack/gommon/log"
	"google.golang.org/appengine"
)

func mobileMenu(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	log.Infof(ctx, "test")

	t, err := template.ParseFiles("../templates/post.html")

	t = t.Funcs(template.FuncMap{"icon": format.IconTag})

	if err != nil {
		w.Write([]byte("Unable to find template"))
	} else {
		t.Execute(w, "Mobile Menu")
	}

	//ctx  = = appengine.NewContext(r)
	// q  = = datastore.NewQuery("Greeting").Ancestor(guestbookKey(ctx)).Order("-Date").Limit(10)
}
