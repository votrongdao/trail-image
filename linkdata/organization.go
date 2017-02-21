package linkdata

type (
	Organization struct {
		Thing
		Alumni    []*Person `json:"alumni"`
		FaxNumber string    `json:"faxNumber"`
		LegalName string    `json:"legalName"`
		Logo      *Image    `json:"logo"`
		NAICS     string    `json:"naics"`
		Founder   []*Person `json:"founder"`
		TaxID     string    `json:"taxID"`
		Telephone string    `json:"telephone"`
	}

	Brand struct {
		Thing
		Ratings *AggregrateRating `json:"aggregateRating"`
		Review  *Review           `json:"review"`
		Logo    *Image            `json:"logo"`
	}
)
