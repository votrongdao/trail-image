package flickr

type (
	// http://www.flickr.com/services/api/response.json.html
	Response struct {
		SetForPhoto *SetForPhoto  `json:"set"`
		SetInfo     *SetInfo      `json:"photoset"`
		SetPhotos   *SetPhotos    `json:"photoset"`
		Status      string        `json:"stat"`
		Code        int           `json:"code"`
		Message     string        `json:"message"`
		Collections []*Collection `json:"collections"`
		Photo       *PhotoMatch   `json:"photo"`
		Sizes       *SizeList     `json:"sizes"`
		PhotoMatch  *PhotoSearch  `json:"photos"`
		TagMatch    *TagSearch    `json:"who"`
	}

	FailResponse struct {
		Stat    string
		Code    int
		Message string
	}
)

func (r *Response) Okay() bool {
	return r.Status == "ok"
}
