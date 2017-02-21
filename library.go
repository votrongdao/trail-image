package trailimage

type Library struct {
	Tags  []string
	Posts []*Post
}

func (l *Library) GetPostWithPhoto(p *Photo) *Post {
	return nil
}

func (l *Library) Add(p *Post) {
	if !l.Contains(p) {
		l.Posts = append(l.Posts, p)
	}
}

func (l *Library) Contains(p *Post) bool {
	for _, v := range l.Posts {
		if v == p {
			return true
		}
	}
	return false
}

// correlatePosts matches posts that are part of a series.
func (l *Library) correlatePosts() {

}
