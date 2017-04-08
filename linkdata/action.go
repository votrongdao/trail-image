package linkdata

type (
	EntryPoint struct {
		Thing
		Application  SoftwareApplication `json:"actionApplication,omitempty"`
		HttpMethod   string              `json:"httpMethod,omitempty"`
		ContentType  string              `json:"contentType,omitempty"`
		EncodingType string              `json:"encodingType,omitempty"`
		UrlTemplate  string              `json:"urlTemplate,omitempty"`
	}

	// See https://schema.org/docs/actions.html
	Action struct {
		Thing
		Target EntryPoint `json:"target,omitempty"`
		// Instrument is the object that helped the agent perform the action.
		// Example: John wrote a book with a pen.
		Instrument Thing `json:"instrument,omitempty"`
	}

	SearchAction struct {
		Action
		Query string `json:"query,omitempty"`
	}
)

// Example
// "potentialAction": {
//      "target": "http://www.trailimage.com/search?q={search_term_string}",
//      "query-input": "required name=search_term_string",
//      "@type": "SearchAction"
//  },
func MakeSearchAction(urlPattern string) *Action {
	return &Action{
		Thing: Thing{Type: "SearchAction"},
	}
}
