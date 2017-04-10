package config

import (
	"os"
)

type (
	drive struct {
		ApiKey       string
		TracksFolder string
	}

	// http://code.google.com/apis/console/#project:1033232213688
	// http://developers.google.com/maps/documentation/staticmaps/
	google struct {
		ApiKey    string
		ProjectID string
		// shown as 'UA-<AnalyticsID>-1
		AnalysticsID   string
		SearchEngineID string
		BlogID         string
		Drive          drive
		Auth           auth
	}
)

var Google = google{
	ApiKey:         os.Getenv("GOOGLE_KEY"),
	ProjectID:      "1033232213688",
	AnalysticsID:   "22180727",
	SearchEngineID: os.Getenv("GOOGLE_SEARCH_ID"),
	BlogID:         "118459106898417641",
	Drive: drive{
		ApiKey:       os.Getenv("GOOGLE_DRIVE_KEY"),
		TracksFolder: "0B0lgcM9JCuSbMWluNjE4LVJtZWM",
	},
	Auth: auth{
		ClientID: os.Getenv("GOOGLE_CLIENT_ID"),
		Secret:   os.Getenv("GOOGLE_SECRET"),
		Callback: "http://www." + DOMAIN + "/auth/google",
		Token: token{
			Access:  os.Getenv("GOOGLE_ACCESS_TOKEN"),
			Refresh: os.Getenv("GOOGLE_REFRESH_TOKEN"),
		},
	},
}
