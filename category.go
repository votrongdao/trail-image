package main

import "strings"

type Category struct {
	Title         string
	Key           string
	Subcategories []*Category
	Posts         []*Post
}

func (c *Category) Add(child *Category) *Category {
	if !c.Contains(child) {
		c.Subcategories = append(c.Subcategories, child)
	}
	return c
}

func (c *Category) Contains(child *Category) bool {
	for _, v := range c.Subcategories {
		if v == child {
			return true
		}
	}
	return false
}

func (c *Category) IsChild() bool {
	return strings.Contains(c.Key, "/")
}

func (c *Category) IsParent() bool {
	return len(c.Subcategories) > 0
}

func (c *Category) RemovePost(p *Post) *Category {
	return c
}
