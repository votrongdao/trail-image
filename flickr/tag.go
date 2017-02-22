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

	TagSearch struct {
		UserID string `json:"id"`
		Tags   struct {
			Tag []*Tag `json:"tag"`
		} `json:"tags"`
	}
)

func (t *Tag) Name() string {
	return t.Original[0].Text
}

func (s *TagSearch) Matches() []*Tag {
	return s.Tags.Tag
}
