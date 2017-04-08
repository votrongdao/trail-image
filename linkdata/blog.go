package linkdata

type (
	// http://schema.org/SocialMediaPosting
	SocialMediaPosting struct {
		Article
		Content *CreativeWork `json:"sharedContent,omitempty"`
	}

	// http://schema.org/BlogPosting
	BlogPosting struct {
		SocialMediaPosting
	}

	// http://schema.org/Blog
	Blog struct {
		CreativeWork
		Posts []*BlogPosting `json:"blogPost"`
	}
)

func MakeBlog(name string) *Blog {
	return &Blog{
		CreativeWork: CreativeWork{
			Thing: Thing{Type: "Blog", Name: name},
		},
	}
}
