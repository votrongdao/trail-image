package flickr

import (
	"bytes"
	"encoding/json"
)

type (
	PhotoMatch struct {
		ServerLocation
		Visibility
		Title  string  `json:"title"`
		Owner  string  `json:"owner"`
		Camera string  `json:"camera"`
		EXIF   []*EXIF `json:"EXIF"`
	}

	PhotoSearch struct {
		Page      uint          `json:"page"`
		PageCount uint          `json:"pages"`
		PerPage   string        `json:"perpage"`
		Total     uint          `json:"total"`
		Photos    []*PhotoMatch `json:"photo"`
	}

	// See https://www.flickr.com/services/api/misc.urls.html
	PhotoSummary struct {
		Place
		ServerLocation
		Coordinate
		Title                string  `json:"title"`
		IsPrimary            uint    `json:"isprimary"`
		Tags                 string  `json:"tags"`
		Description          Content `json:"description"`
		DateTaken            string  `json:"datetaken"`
		DateTakenGranularity uint    `json:"datetakengranularity"`
		DateTakenUnknown     uint    `json:"datetakenunknown"`
		Context              int     `json:"context"`
		geo_is_family        uint
		geo_is_friend        uint
		geo_is_contact       uint
		geo_is_public        uint
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
		Text string
	}

	Owner struct {
		ID         string `json:"nsid"`
		Username   string `json:"username"`
		Location   string `json:"location"`
		IconServer string `json:"iconserver"`
		IconFarm   int    `json:"iconfarm"`
	}

	PhotoMembership struct {
		Set []*SetForPhoto `json:"set"`
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

	EXIF struct {
		TagSpace   string  `json:"tagspace"`
		TagSpaceID uint    `json:"tagspaceid"`
		Tag        string  `json:"tag"`
		Label      string  `json:"label"`
		Raw        Content `json:"raw"`
		Formatted  Content `json:"clean"`
	}
)

func (c *Content) UnmarshalJSON(b []byte) error {
	temp := struct {
		Value json.RawMessage `json:"_content"`
	}{}
	err := json.Unmarshal(b, &temp)
	if err != nil {
		return err
	}
	c.Text = string(bytes.Trim(temp.Value, "\""))

	return nil
}
