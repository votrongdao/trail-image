package linkdata

type (
	Rating struct {
		Thing
		Author *Person `json:"author"`
		Best   string  `json:"bestRating"`
		Value  string  `json:"ratingValue"`
		Worst  string  `json:"worstRating"`
	}

	AggregrateRating struct {
		Rating
		ItemReviews []*Thing `json:"itemReviews"`
		RatingCount uint     `json:"ratingCount"`
		ReviewCount uint     `json:"reviewCount"`
	}

	Review struct {
		Thing
		ItemReviewed *Thing  `json:"itemReviewed"`
		Body         string  `json:"reviewBody"`
		Rating       *Rating `json:"reviewRating"`
	}
)
