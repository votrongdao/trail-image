package flickr

const (
	TYPE_USER  = "user_id"
	TYPE_SET   = "photoset_id"
	TYPE_PHOTO = "photo_id"
)

const (
	URL_HOST          = "api.flickr.com"
	URL_BASE          = "/services/rest/"
	URL_TOKEN_REQUEST = "http://www.flickr.com/services/oauth/request_token"
	URL_AUTHORIZE     = "http://www.flickr.com/services/oauth/authorize"
	URL_TOKEN_ACCESS  = "http://www.flickr.com/services/oauth/access_token"
	URL_PHOTO_SET     = "http://www.flickr.com/photos/trailimage/sets/"
)

const (
	METHOD_COLLECTIONS  = "collections.getTree"
	METHOD_PHOTO_EXIF   = "photos.getExif"
	METHOD_PHOTO_SEARCH = "photos.search"
	METHOD_PHOTO_SETS   = "photos.getAllContexts"
	METHOD_PHOTO_SIZES  = "photos.getSizes"
	METHOD_PHOTO_TAGS   = "tags.getListUserRaw"
	METHOD_SET_INFO     = "photosets.getInfo"
	METHOD_SET_PHOTOS   = "photosets.getPhotos"
)

const (
	EXTRA_DESCRIPTION = "description"
	EXTRA_TAGS        = "tags"
	EXTRA_DATE_TAKEN  = "date_taken"
	EXTRA_LOCATION    = "geo"
	EXTRA_PATH_ALIAS  = "path_alias"
)
