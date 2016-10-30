package model

import (
	"time"
)

type Post struct {
	Author string
	Title string
	Text string
	Tag []string
	Date time.Time
}

func New(Author string, Title string, Text string, Tag []string, Date time.Time) *Post {
	p := &Post{}
	p.Author = Author
	p.Title = Title
	p.Text = Text
	p.Tag = Tag
	p.Date = Date

	return p
}
