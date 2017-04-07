package flickr

import "encoding/json"

type (
	Collection struct {
		ID          string           `json:"id"`
		Title       string           `json:"title"`
		Description string           `json:"description"`
		IconLarge   string           `json:"iconlarge"`
		IconSmall   string           `json:"iconsmall"`
		Collections []*Collection    `json:"collection"`
		Sets        []*CollectionSet `json:"set"`
	}

	CollectionSet struct {
		ID          string `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
	}
)

func ParseCollection(raw []byte) (*Collection, error) {
	out := &Collection{}
	err := json.Unmarshal(raw, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
