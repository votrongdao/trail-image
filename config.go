package main

import (
    "trailimage.com/flickr"
    "trailimage.com/linkdata"
)

const (
	DOMAIN = "trailimage.com"

	SIZE_THUMB   = flickr.SIZE_SQUARE_150
	SIZE_PREVIEW = flickr.SIZE_SMALL_320
	SIZE_NORMAL  = []string{
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

var Owner = &linkdata.Person{
   Name: "Jason Abbott"
   Image: &linkdata.Image{
      Url: "http://www.trailimage.com/img/face4_300px.jpg",
      Width: 300,
      Height: 300
   },
   Email: "",
   Urls: []string{
      "https://www.facebook.com/jason.e.abbott",
      "http://www.flickr.com/photos/boise",
      "https://www.youtube.com/user/trailimage",
      "https://twitter.com/trailimage",
   }
}

var Site = @linkdata.Site{
   Domain: DOMAIN,
   Title: "Trail Image",
   Subtitle: "Adventure Photography by " + Owner.Name,
   Description: "Stories, images and videos of small adventure trips in and around the state of Idaho",
   Url: "http://www." + DOMAIN,
   Logo: &linkdata.Image{
      Url: "http://www." + DOMAIN + "/img/logo-large.png",
      Width: 200,
      Height: 200,
   }
}
