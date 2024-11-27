package models

type Songs struct {
	ID           int    `json:"id"`
	Release_data string `json:"release_data"`
	Group        string `json:"group"`
	Song         string `json:"song"`
	TextParts    string `json:"text_parts"`
	Link         string `json:"link"`
}
