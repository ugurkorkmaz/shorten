package database

import "gorm.io/gorm"

type Shorten struct {
	gorm.Model
	Long string `json:"long"`
	Uid  string `json:"uid"`
}
