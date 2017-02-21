package linkdata

type SoftwareApplication struct {
	Thing
	Category        string `json:"applicationCategory"`
	Suite           string `json:"applicationSuite"`
	DownloadUrl     string `json:"downloadUrl"`
	OperatingSystem string `json:"operatingSystem"`
	Version         string `json:"softwareVersion"`
}
