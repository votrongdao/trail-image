package config

// https://developers.facebook.com/docs/reference/plugins/like/
// https://developers.facebook.com/apps/110860435668134/summary
type facebook struct {
	AppID     string
	PageID    string
	SiteID    string
	AdminID   string
	Enabled   bool
	AuthorURL string
}

var Facebook = facebook{
	AppID:     "110860435668134",
	PageID:    "241863632579825",
	SiteID:    "578261855525416",
	AdminID:   "1332883594",
	Enabled:   true,
	AuthorURL: "https://www.facebook.com/jason.e.abbott",
}
