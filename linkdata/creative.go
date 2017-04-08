package linkdata

type (
	// See http://schema.org/CreativeWork
	CreativeWork struct {
		Thing
		Author            *Person         `json:"author,omitempty"`
		Creator           *Person         `json:"creator,omitempty"`
		Editor            *Person         `json:"editor,omitempty"`
		Height            uint            `json:"height,omitempty"`
		Width             uint            `json:"width,omitempty"`
		Bitrate           string          `json:"bitrate,omitempty"`
		EncodingFormat    string          `json:"encodingFormat,omitempty"`
		ContentUrl        string          `json:"contentUrl,omitempty"`
		CopyrightHolder   *Person         `json:"copyrightHolder,omitempty"`
		CopyrightYear     uint            `json:"copyrightYear,omitempty"`
		Keywords          []string        `json:"keywords,omitempty"`
		IsPartOf          *CreativeWork   `json:"isPartOf,omitempty"`
		HasPart           []*CreativeWork `json:"hasPart,omitempty"`
		Headline          string          `json:"headline,omitempty"`
		Version           float32         `json:"version,omitempty"`
		ProductionCompany *Organization   `json:"productionCompany,omitempty"`
	}

	MediaObject struct {
		CreativeWork
		ContentSize         string `json:"contentSize,omitempty"`
		RequireSubscription bool   `json:"requireSubscription,omitempty"`
	}

	// http://schema.org/ImageObject
	Image struct {
		MediaObject
		Caption   string `json:"caption,omitempty"`
		Thumbnail *Image `json:"thumbnail,omitempty"`
	}

	Video struct {
		MediaObject
		Actor      *Person     `json:"actor,omitempty"`
		Caption    string      `json:"caption,omitempty"`
		Director   *Person     `json:"director,omitempty"`
		MusicBy    *MusicGroup `json:"musicBy,omitempty"`
		Thumbnail  *Image      `json:"thumbnail,omitempty"`
		Transcript string      `json:"transcript,omitempty"`
		FrameSize  string      `json:"videoFrameSize,omitempty"`
		Quality    string      `json:"videoQuality,omitempty"`
	}

	// http://schema.org/Article
	Article struct {
		CreativeWork
		Body       string `json:"articleBody,omitempty"`
		Section    string `json:"articleSection,omitempty"`
		PageStart  uint   `json:"pageStart,omitempty"`
		PageEnd    uint   `json:"pageEnd,omitempty"`
		Pagination string `json:"pagination,omitempty"`
		WordCount  uint   `json:"wordCount,omitempty"`
	}

	NewsArticle struct {
		Article
		Column    string `json:"printColumn,omitempty"`
		Edition   string `json:"printEdition,omitempty"`
		Page      string `json:"printPage,omitempty"`
		Selection string `json:"printSelection,omitempty"`
	}
)

func MakeImage(url string, width uint, height uint) *Image {
	return &Image{
		MediaObject: MediaObject{
			CreativeWork: CreativeWork{
				Width:  width,
				Height: height,
				Thing:  Thing{Type: "ImageObject", URL: url},
			},
		},
	}
}
