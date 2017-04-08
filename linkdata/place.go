package linkdata

type Place struct {
	Thing
	ContainedInPlace *Place `json:"containedInPlace,omitempty"`
}
