package flickr

const (
	TYPE_USER  = "user_id"
	TYPE_SET   = "photoset_id"
	TYPE_PHOTO = "photo_id"
)

const (
	URL_HOST          = "api.flickr.com"
	URL_BASE          = "/services/rest/"
	URL_TOKEN_REQUEST = "http=//www.flickr.com/services/oauth/request_token"
	URL_AUTHORIZE     = "http=//www.flickr.com/services/oauth/authorize"
	URL_TOKEN_ACCESS  = "http=//www.flickr.com/services/oauth/access_token"
	URL_PHOTO_SET     = "http=//www.flickr.com/photos/trailimage/sets/"
)

const (
	EXTRA_DESCRIPTION = "description"
	EXTRA_TAGS        = "tags"
	EXTRA_DATE_TAKEN  = "date_taken"
	EXTRA_LOCATION    = "geo"
	EXTRA_PATH_ALIAS  = "path_alias"
)

const (
	SIZE_THUMB      = "url_t"
	SIZE_SQUARE_75  = "url_sq"
	SIZE_SQUARE_150 = "url_q"
	SIZE_SMALL_240  = "url_s"
	SIZE_SMALL_320  = "url_n"
	SIZE_MEDIUM_500 = "url_m"
	SIZE_MEDIUM_640 = "url_z"
	SIZE_MEDIUM_800 = "url_c"
	SIZE_LARGE_1024 = "url_l"
	SIZE_LARGE_1600 = "url_h"
	SIZE_LARGE_2048 = "url_k"
	SIZE_ORIGINAL   = "url_o"
)
