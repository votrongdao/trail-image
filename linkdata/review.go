package linkdata

type (
	Rating struct {
		Thing
		Author *Person `json:"author,omitempty"`
		Best   string  `json:"bestRating,omitempty"`
		Value  string  `json:"ratingValue,omitempty"`
		Worst  string  `json:"worstRating,omitempty"`
	}

	AggregrateRating struct {
		Rating
		ItemReviews []*Thing `json:"itemReviews,omitempty"`
		RatingCount uint     `json:"ratingCount,omitempty"`
		ReviewCount uint     `json:"reviewCount,omitempty"`
	}

	Review struct {
		Thing
		ItemReviewed *Thing  `json:"itemReviewed,omitempty"`
		Body         string  `json:"reviewBody,omitempty"`
		Rating       *Rating `json:"reviewRating,omitempty"`
	}
)
