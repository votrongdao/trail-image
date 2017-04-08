package linkdata

type (
	Organization struct {
		Thing
		Alumni    []*Person `json:"alumni,omitempty"`
		FaxNumber string    `json:"faxNumber,omitempty"`
		LegalName string    `json:"legalName,omitempty"`
		Logo      *Image    `json:"logo,omitempty"`
		NAICS     string    `json:"naics,omitempty"`
		Founder   []*Person `json:"founder,omitempty"`
		TaxID     string    `json:"taxID,omitempty"`
		Telephone string    `json:"telephone,omitempty"`
	}

	Brand struct {
		Thing
		Ratings *AggregrateRating `json:"aggregateRating,omitempty"`
		Review  *Review           `json:"review,omitempty"`
		Logo    *Image            `json:"logo,omitempty"`
	}
)

func MakeOrganization(name string) *Organization {
	return &Organization{
		Thing: Thing{Type: "Organization", Name: name},
	}
}

// AddLogo adds ImageObject as organization logo.
func (o *Organization) AddLogo(url string, width uint, height uint) *Organization {
	o.Logo = MakeImage(url, width, height)
	return o
}
