package model

// Thumbnail Model
type Thumbnail struct {
	Default string `json:"default"`
	Medium  string `json:"medium"`
	High    string `json:"high"`
}