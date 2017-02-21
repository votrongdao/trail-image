package flickr

type (
	SetSummary struct {
		ID          string  `json:"id"`
		Title       string  `json:"title"`
		Description Content `json:"description"`
	}

	SetInfo struct {
		SetSummary
		ServerLocation
		Owner        string `json:"owner"`
		Username     string `json:"username"`
		Primary      string `json:"primary"`
		Photos       uint   `json:"photos"`
		ViewCount    uint   `json:"count_views"`
		CommentCount uint   `json:"count_comments"`
		PhotoCount   uint   `json:"count_photos"`
		VideoCount   uint   `json:"count_videos"`
		CanComment   bool   `json:"can_comment"`
		DateCreate   uint   `json:"date_create"`
		DateUpdate   uint   `json:"date_update"`
	}

	SetPhotos struct {
		Set       *SetInfo        `json:"photoset"`
		Photos    []*PhotoSummary `json:"photo"`
		Page      uint            `json:"page"`
		PerPage   uint            `json:"per_page"`
		PageCount uint            `json:"pages"`
		Primary   string          `json:"primary"`
		OwnerID   string          `json:"owner"`
		OwnerName string          `json:"ownername"`
		Title     string          `json:"title"`
		Total     uint            `json:"total"`
	}

	SetForPhoto struct {
		ServerLocation
		Title        string `json:"title"`
		Primary      string `json:"primary"`
		ViewCount    uint   `json:"view_count"`
		CommentCount uint   `json:"comment_count"`
		PhotoCount   uint   `json:"count_photo"`
		VideoCount   uint   `json:"count_video"`
	}
)
