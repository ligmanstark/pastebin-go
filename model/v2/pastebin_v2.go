package model_v2

import "time"

type ImagePastebin struct {
	ID        int       `json:"id"`
	UrlSlug   string    `json:"url_slug"`
	CreatedAt time.Time `json:"created_at"`
	ImageData []byte    `json:"image_data"`
	FileSize  int       `json:"file_size"`
	MimeType  string    `json:"mime_type"`
}

type TextPastebin struct {
	ID        int       `json:"id"`
	UrlSlug   string    `json:"url_slug"`
	CreatedAt time.Time `json:"created_at"`
	Content   string    `json:"content,omitempty"`
}
