package middleware

import "net/http"

// Get returns Not Found for any non-GET requests.
func Get(h http.HandlerFunc) http.HandlerFunc {
	return handleMethod(http.MethodGet, h)
}

// Post returns Not Found for any non-POST requests.
func Post(h http.HandlerFunc) http.HandlerFunc {
	return handleMethod(http.MethodPost, h)
}

func handleMethod(method string, handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			http.NotFound(w, r)
		} else {
			handler(w, r)
		}
	}
}
