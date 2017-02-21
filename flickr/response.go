package flickr

type (
	Response struct {
		SetForPhoto *SetForPhoto  `json:"set"`
		SetInfo     []*SetInfo    `json:"photoset"`
		SetPhotos   []*SetPhotos  `json:"photoset"`
		Status      string        `json:"stat"`
		Code        int           `json:"code"`
		Message     string        `json:"message"`
		Collections []*Collection `json:"collections"`
		Photo       *PhotoInfo    `json:"photo"`
		Sizes       []*Size       `json:"sizes"`
		Photos      struct {
			Photo SearchResult `json:"photo"`
		} `json:"photos"`
		Who struct {
			Tags struct {
				Tag []Tag `json:"tag"`
			} `json:"tags"`
		} `json:"who"`
	}

	SearchResult struct {
		Page    uint `json:"page"`
		Pages   uint `json:"pages"`
		PerPage uint `json:"perpage"`
		Total   uint `json:"total"`
	}

	FailResponse struct {
		Stat    string
		Code    int
		Message string
	}
)
