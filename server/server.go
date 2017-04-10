package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

const (
	PATH_CATEGORY      = "category"
	PATH_MONTH         = "month"
	PATH_PART_KEY      = "partKey"
	PATH_PHOTO_ID      = "photoID"
	PATH_PHOTO_TAG     = "tagSlug"
	PATH_POST_ID       = "postID"
	PATH_POST_KEY      = "postKey"
	PATH_ROOT_CATEGORY = "rootCategory"
	PATH_SERIES_KEY    = "seriesKey"
	PATH_YEAR          = "year"
	SLUG               = `:[\w\d-]{4,}`
	TOKEN_PHOTO_ID     = `{` + PATH_PHOTO_ID + `:\d{10,11}}`
	TOKEN_POST_ID      = `{` + PATH_POST_ID + `:\d{17}}`
	TOKEN_POST_KEY     = `{` + PATH_POST_KEY + SLUG + `}`
	TOKEN_SERIES       = `{` + PATH_SERIES_KEY + SLUG + `}/{` + PATH_PART_KEY + SLUG + `}`
)

var r *mux.Router = createRouter()

func createRouter() *mux.Router {
	router := mux.NewRouter()
	http.Handle("/", router)
	return router
}

func init() {
	get := r.Methods(http.MethodGet).Subrouter()
	get.HandleFunc("/", test)
	get.HandleFunc("/rss", test)
	get.HandleFunc("/about", test)
	get.HandleFunc("/js/post-menu-data.js", test)
	get.HandleFunc("/sitemap.xml", test)
	get.HandleFunc("/exif/"+TOKEN_PHOTO_ID, test)
	get.HandleFunc("/category-menu", test)
	get.HandleFunc("/mobile-menu", mobileMenu)
	get.HandleFunc("/search", test)
	get.HandleFunc("/"+TOKEN_PHOTO_ID, test)
	get.HandleFunc("/"+TOKEN_POST_ID, test)
}

// see https://hackernoon.com/golang-template-1-bcb690165663
// see https://golang.org/pkg/text/template/
func test(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	log.Infof(ctx, "test")

	t, err := template.ParseFiles("../templates/post.html")
	if err != nil {
		w.Write([]byte("Unable to find template"))
	} else {
		t.Execute(w, "Trail Image")
	}

	//ctx  = = appengine.NewContext(r)
	// q  = = datastore.NewQuery("Greeting").Ancestor(guestbookKey(ctx)).Order("-Date").Limit(10)
}
