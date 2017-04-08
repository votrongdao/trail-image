package linkdata

const CONTEXT = "http://schema.org"

// See http://schema.org/Person
type Thing struct {
	ID          string    `json:"@id,omitempty"`
	Type        string    `json:"@type,omitempty"`
	Context     string    `json:"@context,omitempty"`
	Description string    `json:"description,omitempty"`
	URL         string    `json:"url,omitempty"`
	Name        string    `json:"name,omitempty"`
	SameAs      []string  `json:"sameAs,omitempty"`
	Image       *Image    `json:"image,omitempty"`
	Action      []*Action `json:"potentialAction,omitempty"`
}

// AddImage adds image details to the base thing.
func (t *Thing) AddImage(url string, width uint, height uint) *Thing {
	t.Image = MakeImage(url, width, height)
	return t
}

// AddUrl adds an identifying URL to the base thing.
func (t *Thing) AddUrl(url string) *Thing {
	t.URL = url
	return t
}

// AddDescription adds a description to the base thing.
func (t *Thing) AddDescription(d string) *Thing {
	t.Description = d
	return t
}

// AddSameAs adds one or more URLs that identify the same thing.
func (t *Thing) AddSameAs(url ...string) *Thing {
	for _, u := range url {
		t.SameAs = append(t.SameAs, u)
	}
	return t
}

// AddAction adds potential actions to the thing.
func (t *Thing) AddAction(target string) *Thing {
	t.Action = append(t.Action, &Action{})
	return t
}
