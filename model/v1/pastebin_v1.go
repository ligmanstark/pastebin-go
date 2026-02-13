package model_v1

import "time"

type Pastebin struct {
	ID        int       `json:"id"`
	Content   string    `json:"content,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UrlSlug   string    `json:"url_slug"`
}
