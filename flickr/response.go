package flickr

type (
	Reponse struct {
		Set         MemberSet     `json:"set"`
		Status      string        `json:"stat"`
		Code        int           `json:"code"`
		Message     string        `json:"message"`
		Collections []*Collection `json:"collections"`
		Photo       PhotoInfo     `json:"photo"`
	}

	SearchResult struct {
		Page    uint `json:"page"`
		Pages   uint `json:"pages"`
		PerPage uint `json:"perpage"`
		Total   uint `json:"total"`
	}
)
