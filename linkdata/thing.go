package linkdata

const CONTEXT = "http://schema.org"

type Thing struct {
	ID      string `json:"@id"`
	Type    string `json:"@type"`
	Context string `json:"@context"`
}
