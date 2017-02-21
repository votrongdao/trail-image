package linkdata

type Site struct {
	Domain      string `json:"domain"`
	Title       string `json:"title"`
	Subtitle    string `json:"subtitle"`
	Description string `json:"description"`
	Url         string `json:"url"`
	Logo        Image  `json:"logo"`
}
