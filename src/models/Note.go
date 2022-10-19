package models

import "gorm.io/gorm"

type Note struct {
	gorm.Model
	Id     string `json:"ID"`
	Name   string `json:"title"`
	Author string `json:"author"`
}
