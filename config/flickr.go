package config

import (
	"os"
)

type (
	setSummary struct {
		ID    string
		Title string
	}

	photoSize struct {
		Post   []string
		Map    []string
		Search []string
	}

	flickr struct {
		UserID      string
		AppID       string
		FeatureSets []setSummary
		ExcludeSets []string
		ExcludeTags []string
		MaxRetries  int
		RetryDelay  int
		PhotoSize   photoSize
		Auth        auth
	}
)

var Flickr = flickr{
	UserID:      "60950751@N04",
	AppID:       "72157631007435048",
	ExcludeSets: []string{"72157631638576162"},
	ExcludeTags: []string{"Idaho", "United States of America", "Abbott", "LensTagger", "Boise"},
	MaxRetries:  10,
	RetryDelay:  300,
	PhotoSize:   photoSize{},
	Auth: auth{
		ClientID: os.Getenv("FLICKR_API_KEY"),
		Secret:   os.Getenv("FLICKR_SECRET"),
		Callback: "http://www." + DOMAIN + "/auth/flickr",
		Token: token{
			Access: os.Getenv("FLICKR_ACCESS_TOKEN"),
			Secret: os.Getenv("FLICKR_TOKEN_SECRET"),
		},
	},
}
