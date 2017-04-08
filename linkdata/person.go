package linkdata

// See http://schema.org/Person
type Person struct {
	Thing
	AdditionalName string        `json:"additionalName,omitempty"`
	Affiliation    *Organization `json:"affiliation,omitempty"`
	Email          string        `json:"email,omitempty"`
	FamilyName     string        `json:"familyName,omitempty"`
	GivenName      string        `json:"givenName,omitempty"`
	JobTitle       string        `json:"jobTitle,omitempty"`
	Knows          []*Person     `json:"knows,omitempty"`
	Spouse         *Person       `json:"spouse,omitempty"`
	Children       []*Person     `json:"children,omitempty"`
}

func MakePerson(name string) *Person {
	return &Person{
		Thing: Thing{Type: "Person", Name: name},
	}
}
