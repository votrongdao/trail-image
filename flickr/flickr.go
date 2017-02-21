package flickr

const (
	base         = "/services/rest/"
	requestToken = "http://www.flickr.com/services/oauth/request_token"
	authorize    = "http://www.flickr.com/services/oauth/authorize"
	accessToken  = "http://www.flickr.com/services/oauth/access_token"
	photoSet     = "http://www.flickr.com/photos/trailimage/sets/"
)

var retries = 0

func call(method, idType, id string) error {
	return nil
}

func parameterize(method, idType, id string) string {
	return ""
}

func GetCollections() interface{} {
	return nil
}

func GetSetInfo() interface{} {
	return nil
}

func GetPhotoSizes() interface{} {
	return nil
}

func GetPhotoContext() interface{} {
	return nil
}

func GetExif() interface{} {
	return nil
}

func GetSetPhotos() interface{} {
	return nil
}
