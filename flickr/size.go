package flickr

type (
	Size struct {
		Label  string `json:"label"`
		Width  uint   `json:"width"`
		Height uint   `json:"height"`
		Source string `json:"source"`
		URL    string `json:"url"`
		Media  string `json:"media"`
	}

	SizeList struct {
		Usage
		Size []*Size `json:"size"`
	}
)
