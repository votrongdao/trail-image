package config

type (
	track struct {
		MinimumLength float64
		MinimumPoints int
	}

	privacy struct {
		// Erase tracks around given latitude and longitude
		// (reverse order from Google map listing).
		Center [2]float64
		Miles  int
		Check  bool
	}

	gpx struct {
		Track track
		// Distance a track point must deviate from others to avoid Douglas-Peucker
		// simplification.
		MaxPointDeviationFeet float64
		// Manually adjusted tracks may have infinite speeds between points so throw
		// out anything over a threshold.
		MaxPossibleSpeed int
		Privacy          privacy
		// Whether track GPX files can be downloaded.
		AllowDownload bool
		// Maximum number of photo markers to show on Google static map. Maximum
		// URL length is 2048; 160 is used for base URL; each coordinate needs
		// about 26 characters.
		MaxMarkers int
	}
)

var Map = gpx{
	Track: track{
		MinimumLength: 0.3,
		MinimumPoints: 5,
	},
	Privacy: privacy{
		Miles: 1,
		Check: false,
	},
	MaxPointDeviationFeet: 0.5,
	MaxPossibleSpeed:      150,
	AllowDownload:         true,
	MaxMarkers:            70,
}
