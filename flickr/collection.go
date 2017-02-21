package flickr

type (
	Collection struct {
		ID          string `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		IconLarge   string `json:"iconlarge"`
		IconSmall   string `json:"iconsmall"`
		Collection  []*Collection
		Set         []*SetSummary
	}
)
