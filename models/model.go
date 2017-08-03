package models


import (
    "github.com/jinzhu/gorm"
)

type Model struct {
	gorm.Model  // gorm will automaticly generate ID with primarykey
    Amount          string      `json:"Amount"`
}

func (Model) TableName() string {
    return "models"
}
