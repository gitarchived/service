package db

import "gorm.io/gorm"

type Host struct {
	gorm.Model `json:"-"`
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	URL        string `json:"url"`
	Prefix     string `json:"prefix"`
}

type Repository struct {
	gorm.Model `json:"-"`
	ID         uint   `json:"id"`
	Owner      string `json:"owner"`
	Name       string `json:"name"`
	Host       string `json:"host"`
	Deleted    bool   `json:"deleted"`
	LastCommit string `json:"last_commit"`
}
