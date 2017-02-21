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

	MediaObject struct {
		CreativeWork
		ContentSize         string `json:"contentSize"`
		RequireSubscription bool   `json:"requireSubscription"`
	}

	// http://schema.org/ImageObject
	Image struct {
		MediaObject
		Caption   string `json:"caption"`
		Thumbnail *Image `json:"thumbnail"`
	}

	Video struct {
		MediaObject
		Actor      *Person     `json:"actor"`
		Caption    string      `json:"caption"`
		Director   *Person     `json:"director"`
		MusicBy    *MusicGroup `json:"musicBy"`
		Thumbnail  *Image      `json:"thumbnail"`
		Transcript string      `json:"transcript"`
		FrameSize  string      `json:"videoFrameSize"`
		Quality    string      `json:"videoQuality"`
	}

	// http://schema.org/Article
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

	// http://schema.org/SocialMediaPosting
	SocialMediaPosting struct {
		Article
		Content *CreativeWork `json:"sharedContent"`
	}

	// http://schema.org/BlogPosting
	BlogPosting struct {
		SocialMediaPosting
	}

	// http://schema.org/Blog
	Blog struct {
		Thing
		Posts []*BlogPosting `json:"blogPost"`
	}
)
