package linkdata

type Person struct {
	Thing
	AdditionalName string        `json:"additionalName"`
	Affiliation    *Organization `json:"affiliation"`
	Email          string        `json:"email"`
	FamilyName     string        `json:"familyName"`
	GivenName      string        `json:"givenName"`
	JobTitle       string        `json:"jobTitle"`
	Knows          []*Person     `json:"knows"`
	Spouse         *Person       `json:"spouse"`
}
