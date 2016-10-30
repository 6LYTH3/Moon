package model

import (
	"time"
)

type Post struct {
	Author string
	Title  string
	Text   string
	Tag    []string
	Date   time.Time
}
