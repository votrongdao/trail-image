package flickr

type (
	Usage struct {
		CanDownload bool `json:"candownload"`
		CanBlog     bool `json:"canblog"`
		CanPrint    bool `json:"canprint"`
		CanShare    bool `json:"canshare"`
	}

	Visibility struct {
		IsPublic bool `json:"ispublic"`
		IsFriend bool `json:"isfriend"`
		IsFamily bool `json:"isfamily"`
	}

	Permission struct {
		ToComment     bool `json:"permcomment"`
		ToSeeMetadata bool `json:"permmetadata"`
	}

	EditAbility struct {
		AddComment  bool `json:"cancomment"`
		AddMetadata bool `json:"canaddmeta"`
	}

	LocationPermission struct {
		Visibility
		IsContent bool `json:"iscontent"`
	}
)
