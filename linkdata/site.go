package linkdata

// type Site struct {
// 	Domain      string `json:"domain,omitempty"`
// 	Title       string `json:"title,omitempty"`
// 	Subtitle    string `json:"subtitle,omitempty"`
// 	Description string `json:"description,omitempty"`
// 	Url         string `json:"url,omitempty"`
// 	Logo        Image  `json:"logo,omitempty"`
// }

// See http://schema.org/WebSite
type WebSite struct {
	CreativeWork
	AccountablePerson *Person `json:"accountablePerson,omitempty"`
	Author            *Person `json:"person,omitempty"`
	Editor            *Person `json:"editor,omitempty"`
	ThumbnailUrl      string  `json:"thumbnailUrl,omitempty"`
	Version           int     `json:"version,omitempty"`
}

func MakeWebSite(title string) *WebSite {
	return &WebSite{
		CreativeWork: CreativeWork{
			Thing: Thing{Type: "WebSite", Name: title},
		},
	}
}
