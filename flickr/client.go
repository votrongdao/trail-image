package flickr

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

type (
	Client struct {
		UserID       string
		ApiKey       string
		Token        string
		Secret       string
		Signature    string
		RequestToken string
		AccessToken  string
		CallbackUrl  string
	}

	Params map[string]string
)

// http://www.flickr.com/services/api/response.json.html
func (c *Client) call(method, idType, id string, extras Params) (*Response, error) {
	//key := method + ":" + id
	url := "https://" + URL_HOST + URL_BASE + c.parameterize(method, idType, id, extras)

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	api := &Response{}

	if err = json.Unmarshal(body, api); err != nil {
		return nil, err
	}
	return api, nil
}

// https://github.com/mrjones/oauth/blob/master/examples/netflix/netflix.go
func (c *Client) authCall() error {
	// client := oauth.NewConsumer(
	// 	c.ApiKey,
	// 	c.Secret,
	// 	oauth.ServiceProvider{
	// 		RequestTokenUrl:   URL_TOKEN_REQUEST,
	// 		AuthorizeTokenUrl: URL_AUTHORIZE,
	// 		AccessTokenUrl:    URL_TOKEN_ACCESS,
	// 	})
	return nil
}

func (c *Client) parameterize(method, idType, id string, extras Params) string {
	qs := ""
	op := "?"
	args := Params{
		"api_key":        c.ApiKey,
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

func (c *Client) GetCollections() ([]*Collection, error) {
	res, err := c.call("collections.getTree", TYPE_USER, c.UserID, nil)
	if err != nil {
		return nil, err
	}
	return res.Collections, nil
	// call(method.COLLECTIONS, type.USER, config.flickr.userID
}

func (c *Client) GetSetInfo(setID string) (*SetInfo, error) {
	res, err := c.call("photosets.getInfo", TYPE_SET, setID, nil)
	if err != nil {
		return nil, err
	}
	return res.SetInfo, nil
	//getSetInfo: id => call(method.set.INFO, type.SET, id, { value: r => r.photoset, allowCache: true }),
}

func (c *Client) GetSetPhotos(setID string) (*SetPhotos, error) {
	res, err := c.call("photosets.getPhotos", TYPE_SET, setID, Params{
		"extras": strings.Join([]string{
			EXTRA_DESCRIPTION,
			EXTRA_TAGS,
			EXTRA_DATE_TAKEN,
			EXTRA_LOCATION,
			EXTRA_PATH_ALIAS,
		}, ","),
	})
	if err != nil {
		return nil, err
	}
	return res.SetPhotos, nil
	// call(method.set.PHOTOS, type.SET, id
}

func (c *Client) GetPhotoSizes(photoID string) ([]*Size, error) {
	res, err := c.call("photos.getSizes", TYPE_PHOTO, photoID, nil)
	if err != nil {
		return nil, err
	}
	return res.Sizes.Size, nil
	//call(method.photo.SIZES, type.PHOTO, id, { value: r => r.sizes.size }),
}

func (c *Client) GetTaggedPhotos(tags []string) (*PhotoSearch, error) {
	res, err := c.call("photos.getSizes", TYPE_USER, c.UserID, Params{
		"extras": strings.Join([]string{
			EXTRA_DESCRIPTION,
			EXTRA_TAGS,
			EXTRA_DATE_TAKEN,
			EXTRA_LOCATION,
			EXTRA_PATH_ALIAS,
		}, ","),
		"tags":     strings.Join(tags, ","),
		"sort":     "relevance",
		"per_page": "500",
	})

	if err != nil {
		return nil, err
	}
	return res.PhotoMatch, nil
}

func (c *Client) GetPhotoContext(photoID string) (*SetForPhoto, error) {
	res, err := c.call("photos.getAllContexts", TYPE_PHOTO, photoID, nil)
	if err != nil {
		return nil, err
	}
	return res.SetForPhoto, nil

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
	// call(method.photo.SETS, type.PHOTO, id, { value: r => r.set }),
}

func (c *Client) GetExif(photoID string) ([]*EXIF, error) {
	res, err := c.call("photos.getExif", TYPE_PHOTO, photoID, nil)
	if err != nil {
		return nil, err
	}
	return res.Photo.EXIF, nil
	// call(method.photo.EXIF, type.PHOTO, id, { value: r => r.photo.exif, allowCache: true }),
}

func (c *Client) GetUserTags() ([]*Tag, error) {
	res, err := c.call("tags.getListUserRaw", TYPE_USER, c.UserID, nil)
	if err != nil {
		return nil, err
	}
	return res.TagMatch.Matches(), nil
}
