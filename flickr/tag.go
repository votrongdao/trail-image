package flickr

type (
	Tag struct {
		Slug     string    `json:"clean"`
		Original []Content `json:"raw"`
	}

	TagSummary struct {
		ID         string `json:"id"`
		Author     string `json:"author"`
		Raw        string `json:"raw"`
		MachineTag int    `json:"machine_tag"`
	}
)
