package main

import (
	"trailimage.com/flickr"
	ld "trailimage.com/linkdata"
)

const (
	NAME         = "Trail Image"
	DOMAIN       = "trailimage.com"
	SIZE_THUMB   = flickr.SIZE_SQUARE_150
	SIZE_PREVIEW = flickr.SIZE_SMALL_320
)

var (
	SIZE_NORMAL = []string{
		flickr.SIZE_LARGE_1024,
		flickr.SIZE_MEDIUM_800,
		flickr.SIZE_MEDIUM_640,
	}
	SIZE_BIG = []string{
		flickr.SIZE_LARGE_2048,
		flickr.SIZE_LARGE_1600,
		flickr.SIZE_LARGE_1024,
	}
)

var Owner = ld.MakePerson("Jason Abbott").
	AddImage("http://www.trailimage.com/img/face4_300px.jpg", 300, 300).
	AddUrl("http://www.trailimage.com/about").
	AddSameAs(
		"https://www.facebook.com/jason.e.abbott",
		"http://www.flickr.com/photos/boise",
		"https://www.youtube.com/user/trailimage",
		"https://twitter.com/trailimage",
	)

var Site = ld.MakeWebSite(NAME).
	AddPublisher(ld.
		MakeOrganization(NAME).
		AddLogo("http://www."+DOMAIN+"/img/logo-large.png", 200, 200)).
	AddDescription("Stories, images and videos of small adventure trips in and around the state of Idaho").
	AddUrl("http://www." + DOMAIN)

	// var Sites = &linkdata.Site{
	// 	Domain:      DOMAIN,
	// 	Title:       "Trail Image",
	// 	Subtitle:    "Adventure Photography by " + Owner.Name,
	// 	Description: "Stories, images and videos of small adventure trips in and around the state of Idaho",
	// 	Url:         "http://www." + DOMAIN,
	// 	Logo: &linkdata.Image{
	// 		Url:    "http://www." + DOMAIN + "/img/logo-large.png",
	// 		Width:  200,
	// 		Height: 200,
	// 	},
	// }
