package flickr

type EXIF struct {
	TagSpace   string  `json:"tagspace"`
	TagSpaceID int     `json:"tagspaceid"`
	Tag        string  `json:"tag"`
	Label      string  `json:"label"`
	Raw        Content `json:"raw"`
}
