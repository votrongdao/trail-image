package linkdata

type (
	CreativeWork struct {
		Thing
		Author            *Person         `json:"author"`
		Creator           *Person         `json:"creator"`
		Editor            *Person         `json:"editor"`
		Height            uint            `json:"height"`
		Width             uint            `json:"width"`
		Bitrate           string          `json:"bitrate"`
		EncodingFormat    string          `json:"encodingFormat"`
		ContentUrl        string          `json:"contentUrl"`
		CopyrightHolder   *Person         `json:"copyrightHolder"`
		CopyrightYear     uint            `json:"copyrightYear"`
		Keywords          []string        `json:"keywords"`
		IsPartOf          *CreativeWork   `json:"isPartOf"`
		HasPart           []*CreativeWork `json:"hasPart"`
		Headline          string          `json:"headline"`
		Version           float32         `json:"version"`
		ProductionCompany *Organization   `json:"productionCompany"`
	}

	// http://schema.org/ImageObject
	Image struct {
		Url    string `json:"url"`
		Width  uint   `json:"width"`
		Height uint   `json:"height"`
	}

	Article struct {
		CreativeWork
		Body       string `json:"articleBody"`
		Section    string `json:"articleSection"`
		PageStart  uint   `json:"pageStart"`
		PageEnd    uint   `json:"pageEnd"`
		Pagination string `json:"pagination"`
		WordCount  uint   `json:"wordCount"`
	}

	NewsArticle struct {
		Article
		Column    string `json:"printColumn"`
		Edition   string `json:"printEdition"`
		Page      string `json:"printPage"`
		Selection string `json:"printSelection"`
	}

	SocialMediaPosting struct {
		Article
		Content *CreativeWork `json:"sharedContent"`
	}

	BlogPosting struct {
		SocialMediaPosting
	}

	Blog struct {
		Thing
		Posts []*BlogPosting `json:"blogPost"`
	}
)
