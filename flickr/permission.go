package flickr

type (
	Usage struct {
		CanDownload uint `json:"candownload"`
		CanBlog     uint `json:"canblog"`
		CanPrint    uint `json:"canprint"`
		CanShare    uint `json:"canshare"`
	}

	Visibility struct {
		IsPublic uint `json:"ispublic"`
		IsFriend uint `json:"isfriend"`
		IsFamily uint `json:"isfamily"`
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
