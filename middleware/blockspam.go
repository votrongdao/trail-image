package middleware

import "net/http"
import "time"

var (
	lastUpdate = time.Now()
	blackList  []string
)

func BlockSpamReferer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if blackListed(r.Referer()) {
			http.NotFound(w, r)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

// blackListed indicates if a domain is found in the black list.
func blackListed(domain string) bool {
	if len(blackList) == 0 {
		downloadBlackList()
	}

	for _, d := range blackList {
		if d == domain {
			return true
		}
	}
	return false
}

func downloadBlackList() {

}
