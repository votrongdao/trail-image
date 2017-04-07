package library

import (
	"fmt"
	"time"

	"trailimage.com/flickr"
)

type Post struct {
	ID            string
	Chronological bool
	PhotosLoaded  bool
	InfoLoaded    bool
	TriedTrack    bool
	// whether a GPS track was found for the post
	HasTrack      bool
	originalTitle string
	PhotoCount    uint
	Photos        []*Photo
	CoverPhoto    *Photo

	CreatedOn time.Time
	UpdatedOn time.Time

	Title    string
	SubTitle string

	Description     string
	LongDescription string

	IsPartial bool
	// position of this post in a series
	Part uint
	// whether next post is part of the same series
	NextIsPart bool
	// whether previous post is part of the same series
	PreviousIsPart bool
	TotalParts     uint
	IsSeriesStart  bool

	Key       string
	SeriesKey string
	PartKey   string

	Next     *Post
	Previous *Post

	BigThumbUrl   string
	SmallThumbUrl string
}

func (p *Post) MakeSeriesStart() *Post {
	p.IsSeriesStart = true
	p.Key = p.SeriesKey

	return p
}

func (p *Post) Ungroup() {

}

func (p *Post) HasKey(key string) bool {
	return p.Key == key || (p.PartKey != "" && key == p.SeriesKey+"-"+p.PartKey)
}

// Empty removes all post information.
func (p *Post) Empty() *Post {
	p.Title = p.originalTitle
	p.SubTitle = ""
	p.Part = 0
	p.TotalParts = 0
	p.IsSeriesStart = false
	p.IsPartial = false
	p.NextIsPart = false
	p.PreviousIsPart = false
	p.SeriesKey = ""
	p.PartKey = ""

	return p
}

// RemoveInfo clears all post fields.
func (p *Post) RemoveInfo() *Post {
	//p.CreatedOn = time.Time.
	p.CoverPhoto = nil
	p.LongDescription = ""
	p.PhotoCount = 0

	p.InfoLoaded = false
	p.TriedTrack = false

	// from getPhotos()
	p.Photos = nil
	p.PhotosLoaded = false

	return p
}

// ParseSetInfo converts Flickr response into a Post.
func ParseSetInfo(info *flickr.SetInfo) *Post {
	thumb := fmt.Sprintf("http://farm%d.staticflicker.com/%s/%s_%s", info.Farm, info.Server, info.Primary, info.Secret)
	return &Post{
		ID:            info.ID,
		PhotoCount:    info.PhotoCount,
		Description:   info.Description.Text,
		BigThumbUrl:   thumb + ".jpg",
		SmallThumbUrl: thumb + "_s.jpg",
		InfoLoaded:    true,
	}
}

// const thumb = `http://farm${setInfo.farm}.staticflickr.com/${setInfo.server}/${setInfo.primary}_${setInfo.secret}`;

//       return Object.assign(this, {
//          // removes video information from setInfo.description
//          video: buildVideoInfo(setInfo),
//          createdOn: util.date.fromTimeStamp(setInfo.date_create),
//          updatedOn: util.date.fromTimeStamp(setInfo.date_update),
//          photoCount: setInfo.photos,
//          description: setInfo.description._content.remove(/[\r\n\s]*$/),
//          // long description is updated after photos are loaded
//          longDescription: this.description,
//          // http://farm{farm-id}.staticflickr.com/{server-id}/{id}_{secret}_[mstzb].jpg
//          // http://farm{{info.farm}}.static.flickr.com/{{info.server}}/{{info.primary}}_{{info.secret}}.jpg'
//          // thumb URLs may be needed before photos are loaded, e.g. in RSS XML
//          bigThumbURL: thumb + '.jpg',     // 500px
//          smallThumbURL: thumb + '_s.jpg',
//          infoLoaded: true
//       });
