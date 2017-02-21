package flickr

import (
	"encoding/json"
	"errors"
	"fmt"
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
	base         = "/services/rest/"
	requestToken = "http://www.flickr.com/services/oauth/request_token"
	authorize    = "http://www.flickr.com/services/oauth/authorize"
	accessToken  = "http://www.flickr.com/services/oauth/access_token"
	photoSet     = "http://www.flickr.com/photos/trailimage/sets/"
)

var retries = 0

func parse(data []byte, v interface{}) error {
	fail := Fail(data)

	if fail != nil {
		return fail
	}

	err := json.Unmarshal(data, v)

	if err != nil {
		fail := Fail(data)

		if fail != nil {
			return fail
		}

		return err
	}

	return nil
}

func Fail(data []byte) error {
	fail := &FailResponse{}
	err := json.Unmarshal(data, fail)

	if err == nil && fail.Stat == "fail" {
		return errors.New(fail.Message)
	}

	return nil
}

type Params map[string]string

type Client struct {
	Key   string
	Token string
	Sig   string
}

func (client *Client) Request(method string, params Params) ([]byte, error) {
	url := fmt.Sprintf("https://api.flickr.com/services/rest/?method=flickr.%s&api_key=%s&format=json&nojsoncallback=1", method, client.Key)

	if len(client.Token) > 0 {
		url = fmt.Sprintf("%s&auth_token=%s", url, client.Token)
	}

	if len(client.Sig) > 0 {
		url = fmt.Sprintf("%s&auth_sig=%s", url, client.Sig)
	}

	for key, value := range params {
		url = fmt.Sprintf("%s&%s=%s", url, key, value)
	}

	response, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	return ioutil.ReadAll(response.Body)
}

func call(method, idType, id string) error {
	key := method + ":" + id
	methodUrl := "https://" + host + base + parameterize(method, idType, id, options.args)

	return nil
}

func parameterize(method, idType, id string) string {
	qs := ""
	op := "?"

	args := map[string]string{
		"api_key":        "",
		"format":         "json",
		"nojsoncallback": "1",
		"method":         "flickr." + method,
	}

	if idType != "" && id != "" {
		args[idType] = id
	}

	// let qs = '';
	// let op = '?';

	// args.api_key = config.flickr.auth.apiKey;
	// args.format = 'json';
	// args.nojsoncallback = 1;
	// args.method = 'flickr.' + method;

	// if (is.value(idType) && is.value(id)) { args[idType] = id; }
	// for (const k in args) { qs += op + k + '=' + encodeURIComponent(args[k]); op = '&'; }
	// return qs;

	for k, v := range args {
		qs += op + k + "=" + v
		op = "&"
	}

	return qs
}

func GetCollections(id string) interface{} {
	return call(method.COLLECTIONS, TYPE_USER, id)
	// call(method.COLLECTIONS, type.USER, config.flickr.userID
}

func GetSetInfo(id string) interface{} {
	return call(method.SET_INFO, TYPE_SET, id)
	//getSetInfo: id => call(method.set.INFO, type.SET, id, { value: r => r.photoset, allowCache: true }),
}

func GetPhotoSizes(id string) interface{} {
	return call(method.PHOTO_SIZES, TYPE_PHOTO, id)
	//call(method.photo.SIZES, type.PHOTO, id, { value: r => r.sizes.size }),
}

func GetPhotoContext(id string) interface{} {
	return call(method.PHOTO_SETS, TYPE_PHOTO, id)
	// call(method.photo.SETS, type.PHOTO, id, { value: r => r.set }),
}

func GetExif(id string) interface{} {
	return call(method.PHOTO_EXIF, TYPE_PHOTO, id)
	// call(method.photo.EXIF, type.PHOTO, id, { value: r => r.photo.exif, allowCache: true }),
}

func GetSetPhotos(id string) interface{} {
	return call(method.SET_PHOTOS, TYPE_SET, id)
	// call(method.set.PHOTOS, type.SET, id
}
