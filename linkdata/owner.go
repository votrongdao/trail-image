package linkdata

type Owner struct {
	Name  string   `json:"name"`
	Email string   `json:"email"`
	Urls  []string `json:"urls"`
	Image Image    `json:"image"`
}
