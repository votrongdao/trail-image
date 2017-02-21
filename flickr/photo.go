package flickr

type (
	// See https://www.flickr.com/services/api/misc.urls.html
	PhotoSummary struct {
		Place
		ServerLocation
		Coordinate
		Title                string  `json:"title"`
		IsPrimary            bool    `json:"isprimary"`
		Tags                 string  `json:"tags"`
		Description          Content `json:"description"`
		DateTaken            string  `json:"datetaken"`
		DateTakenGranularity string  `json:"datetakengranularity"`
		Context              int     `json:"context"`
		geo_is_family        bool
		geo_is_friend        bool
		geo_is_contact       bool
		geo_is_public        bool
		LastUpdate           string `json:"lastupdate"`
		PathAlias            string `json:"pathalias"`

		SquareURL    string `json:"url_s"`
		SquareHeight string `json:"height_s"`
		SquareWidth  string `json:"width_s"`

		Large1600URL    string `json:"url_h"`
		Large1600Height string `json:"height_h"`
		Large1600Width  string `json:"width_h"`

		Large2048URL    string `json:"url_k"`
		Large2048Height string `json:"height_k"`
		Large2048Width  string `json:"width_k"`

		SmallURL    string `json:"url_m"`
		SmallHeight string `json:"height_m"`
		SmallWidth  string `json:"width_m"`

		OriginalURL    string `json:"url_o"`
		OriginalHeight string `json:"height_o"`
		OriginalWidth  string `json:"width_o"`
	}

	//    url_l?: string,
	//    height_l?: string,
	//    width_l?: string,

	Content struct {
		Text string `json:"_content"`
	}

	Owner struct {
		ID         string `json:"nsid"`
		Username   string `json:"username"`
		Location   string `json:"location"`
		IconServer string `json:"iconserver"`
		IconFarm   int    `json:"iconfarm"`
	}

	PhotoMembership struct {
		Set []*MemberSet `json:"set"`
	}

	PhotoDates struct {
		Posted           string `json:"posted"`
		Taken            string `json:"taken"`
		TakenGranularity int    `json:"takengranularity"`
		LastUpdate       string `json:"lastupdate"`
	}

	PhotoInfo struct {
		ServerLocation
		DateUploaded      string      `json:"dateuploaded"`
		IsFavorite        bool        `json:"isfavorite"`
		License           string      `json:"license"`
		SafetyLevel       uint        `json:"safetylevel"`
		Rotate            bool        `json:"rotate"`
		OriginalSecret    string      `json:"originalsecret"`
		OriginalFormat    string      `json:"originalformat"`
		Owner             Owner       `json:"owner"`
		Title             Content     `json:"title"`
		Description       Content     `json:"description"`
		Visibility        Visibility  `json:"visibility"`
		Dates             PhotoDates  `json:"dates"`
		Views             uint        `json:"views"`
		Permissions       Permission  `json:"permissions"`
		Editability       EditAbility `json:"editability"`
		PublicEditability EditAbility `json:"publiceditability"`
		Usage             Usage       `json:"usage"`
		Tags              int
		Location          Location           `json:"location"`
		GeoPermission     LocationPermission `json:"geoperms"`
		Media             string             `json:"media"`
		URLs              string
	}
)
