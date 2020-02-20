package model

type Database struct {
	Title string `json:"title" bson:"title,omitempty"`
	Slug  string `json:"slug" bson:"slug,omitempty"`
}
