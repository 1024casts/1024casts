package model

import "time"

type CategoryModel struct {
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	DeletedAt   time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	Description string    `gorm:"column:description" json:"description"`
	Id          int       `gorm:"column:id" json:"id"`
	Name        string    `gorm:"column:name" json:"name"`
	ParentId    int       `gorm:"column:parent_id" json:"parent_id"`
	Slug        string    `gorm:"column:slug" json:"slug"`
	TopicCount  int       `gorm:"column:topic_count" json:"topic_count"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at"`
	UserID      int       `gorm:"column:user_id" json:"user_id"`
	Weight      int       `gorm:"column:weight" json:"weight"`
}

// TableName sets the insert table name for this struct type
func (c *CategoryModel) TableName() string {
	return "forum_categories"
}
