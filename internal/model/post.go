package model

import "time"

type Post struct {
	ID          int       `gorm:"column:id;primaryKey;autoIncrement"`
	Title       string    `gorm:"column:title"`
	Content     string    `gorm:"column:content"`
	Category    string    `gorm:"column:category"`
	CreatedDate time.Time `gorm:"column:created_date"`
	UpdatedDate time.Time `gorm:"column:updated_date"`
	Status      string    `gorm:"column:status"`
}

func (Post) TableName() string {
	return "posts"
}
