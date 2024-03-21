package model

// Video Model 
type Video struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	PublishedAt string     `json:"published_at"`
	Thumbnail  Thumbnail `json:"thumbnail"`
}