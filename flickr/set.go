package flickr

type (
	SetSummary struct {
		ID          string `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	SetInfo struct {
		SetSummary
		Owner    string `json:"owner"`
		Username string `json:"username"`
	}

	MemberSet struct {
		FarmLocation
		Title        string `json:"title"`
		Primary      string `json:"primary"`
		ViewCount    uint   `json:"view_count"`
		CommentCount uint   `json:"comment_count"`
		PhotoCount   uint   `json:"count_photo"`
		VideoCount   uint   `json:"count_video"`
	}
)
