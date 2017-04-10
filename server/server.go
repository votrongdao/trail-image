package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

const (
	ROUTE_CATEGORY      = "category"
	ROUTE_MONTH         = "month"
	ROUTE_PART_KEY      = "partKey"
	ROUTE_PHOTO_ID      = "photoID"
	ROUTE_PHOTO_TAG     = "tagSlug"
	ROUTE_POST_ID       = "postID"
	ROUTE_POST_KEY      = "postKey"
	ROUTE_ROOT_CATEGORY = "rootCategory"
	ROUTE_SERIES_KEY    = "seriesKey"
	ROUTE_YEAR          = "year"
	RE_SLUG             = `([\w\d-]{4,})`
	RE_PHOTO_ID         = `:` + ROUTE_PHOTO_ID + `(\d{10,11})`
	RE_POST_ID          = `:` + ROUTE_POST_ID + `(\d{17})`
	RE_POST_KEY         = `:` + ROUTE_POST_KEY + RE_SLUG
	RE_SERIES           = `:` + ROUTE_SERIES_KEY + RE_SLUG + `/:` + ROUTE_PART_KEY + RE_SLUG
)

var r = createRouter()

func createRouter() *mux.Router {
	router := mux.NewRouter()
	http.Handle("/", router)
	return router
}

func init() {
	r.HandleFunc("/", test).Methods(http.MethodGet)
	r.HandleFunc("/products", test)
	r.HandleFunc("/articles", test)
}

// see https = //hackernoon.com/golang-template-1-bcb690165663
func test(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	log.Infof(ctx, "test")

	t, err := template.ParseFiles("../templates/post.html")
	if err != nil {
		w.Write([]byte("Unable to find template"))
	} else {
		t.Execute(w, "Hello World!")
	}

	//ctx  = = appengine.NewContext(r)
	// q  = = datastore.NewQuery("Greeting").Ancestor(guestbookKey(ctx)).Order("-Date").Limit(10)
}
