package trailimage

import "time"

type Post struct {
	ID            string
	Chronological bool
	PhotosLoaded  bool
	InfoLoaded    bool
	TriedTrack    bool
	// HasTrack indicates whether a GPS track was found for the post
	HasTrack      bool
	originalTitle string
	PhotoCount    int
	Photos        []*Photo
	CoverPhoto    *Photo

	CreatedOn time.Time
	UpdatedOn time.Time

	Title    string
	SubTitle string

	LongDescription string

	IsPartial bool
	// position of this post in a series
	Part int
	// whether next post is part of the same series
	NextIsPart bool
	// whether previous post is part of the same series
	PreviousIsPart bool
	TotalParts     int
	IsSeriesStart  bool

	Key       string
	SeriesKey string
	PartKey   string

	Next     *Post
	Previous *Post
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
