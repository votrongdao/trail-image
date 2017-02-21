package linkdata

type Action struct {
	Thing
	Instrument Thing `json:"instrument"`
}
