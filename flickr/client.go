package flickr

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"trailimage.com/flickr/method"
)

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

type (
	Client struct {
		Key       string
		Token     string
		Signature string
	}

	Params map[string]string
)

func call(method, idType, id string, out interface{}, extras Params) error {
	key := method + ":" + id
	url := "https://" + URL_HOST + URL_BASE + parameterize(method, idType, id, extras)

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, out)
}

func parameterize(method, idType, id string, extras Params) string {
	qs := ""
	op := "?"

	args := Params{
		"api_key":        "",
		"format":         "json",
		"nojsoncallback": "1",
		"method":         "flickr." + method,
	}

	if idType != "" && id != "" {
		args[idType] = id
	}

	if extras != nil {
		for k, v := range extras {
			args[k] = v
		}
	}

	for k, v := range args {
		// TODO: encodeURIComponent
		qs += op + k + "=" + v
		op = "&"
	}

	return qs
}

func GetCollections(id string) (*Collection, error) {
	out := &Collection{}
	if err := call(method.COLLECTIONS, TYPE_USER, id, out, nil); err != nil {
		return nil, err
	}
	return out, nil
	// call(method.COLLECTIONS, type.USER, config.flickr.userID
}

func GetSetInfo(id string) (*SetInfo, error) {
	out := &SetInfo{}
	if err := call(method.SET_INFO, TYPE_SET, id, out, nil); err != nil {
		return nil, err
	}
	return out, nil
	//getSetInfo: id => call(method.set.INFO, type.SET, id, { value: r => r.photoset, allowCache: true }),
}

func GetPhotoSizes(id string) interface{} {
	res, err := call(method.PHOTO_SIZES, TYPE_PHOTO, id, nil)
	if err != nil {
		return nil, err
	}
	return nil, nil
	//call(method.photo.SIZES, type.PHOTO, id, { value: r => r.sizes.size }),
}

func GetPhotoContext(id string) (*PhotoInfo, error) {
	empty := &PhotoInfo{}
	if err := call(method.PHOTO_SETS, TYPE_PHOTO, id, nil); err != nil {
		return nil, err
	}
	return empty, nil

	// info := &struct {
	// 	PhotoSet struct {
	// 		Title struct {
	// 			Content string `json:"_content"`
	// 		} `json:"title"`
	// 		Description struct {
	// 			Content string `json:"_content"`
	// 		} `json:"description"`
	// 	} `json:"photoset"`
	// }{}

	if err := Parse(res, info); err != nil {
		return nil, err
	}
	return nil, nil
	// call(method.photo.SETS, type.PHOTO, id, { value: r => r.set }),
}

func GetExif(id string) interface{} {
	res, err := call(method.PHOTO_EXIF, TYPE_PHOTO, id, nil)
	if err != nil {
		return nil, err
	}
	return nil, nil
	// call(method.photo.EXIF, type.PHOTO, id, { value: r => r.photo.exif, allowCache: true }),
}

func GetSetPhotos(id string) (*SetInfo, error) {
	res, err := call(method.SET_PHOTOS, TYPE_SET, id, nil)
	if err != nil {
		return nil, err
	}
	return nil, nil
	// call(method.set.PHOTOS, type.SET, id
}
