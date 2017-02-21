package flickr

type (
	PhotoSummary struct {
		Place
		ID string `json:"id"`
	}

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
		FarmLocation
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
