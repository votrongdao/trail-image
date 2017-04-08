package linkdata

type SoftwareApplication struct {
	Thing
	Category        string `json:"applicationCategory,omitempty"`
	Suite           string `json:"applicationSuite,omitempty"`
	DownloadUrl     string `json:"downloadUrl,omitempty"`
	OperatingSystem string `json:"operatingSystem,omitempty"`
	Version         string `json:"softwareVersion,omitempty"`
}
