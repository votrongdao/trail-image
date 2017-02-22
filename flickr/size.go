package flickr

import (
	"bytes"
	"encoding/json"
	"strconv"
)

type (
	Size struct {
		Label  string
		Width  uint
		Height uint
		Source string
		URL    string
		Media  string
	}

	SizeList struct {
		Usage
		Size []*Size `json:"size"`
	}
)

func toUint(m json.RawMessage) uint {
	i, err := strconv.ParseInt(string(bytes.Trim(m, "\"")), 10, 32)
	if err != nil {
		return 0
	}
	return uint(i)
}

func (s *Size) UnmarshalJSON(b []byte) error {
	temp := struct {
		Label  string          `json:"label"`
		Source string          `json:"source"`
		Width  json.RawMessage `json:"width"`
		Height json.RawMessage `json:"height"`
		URL    string          `json:"url"`
		Media  string          `json:"media"`
	}{}
	err := json.Unmarshal(b, &temp)
	if err != nil {
		return err
	}
	s.Label = temp.Label
	s.Source = temp.Source
	s.URL = temp.URL
	s.Media = temp.Media
	s.Width = toUint(temp.Width)
	s.Height = toUint(temp.Height)

	return nil
}
