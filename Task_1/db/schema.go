package db

import "gorm.io/gorm"

type Book struct {
	gorm.Model        // this is a struct embedded in Book struct which has fields like ID, CreatedAt, UpdatedAt, DeletedAt
	Title      string `json:"title"`
	Author     string `json:"author"`
}

type Movie struct {
	gorm.Model
	Title  string `json:"title"`
	Rating int    `json:"rating"`
}

type Music struct {
	gorm.Model
	Title  string `json:"title"`
	Artist string `json:"artist"`
}
