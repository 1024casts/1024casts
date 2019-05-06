package model

import (
	"sync"
)

// User represents a registered user.
type CourseModel struct {
	BaseModel
	Name         string `json:"name" gorm:"column:name;not null" binding:"required" validate:"min=1,max=32"`
	Type         string `json:"type" gorm:"column:type;not null"`
	Description  string `json:"description" gorm:"column:description;" binding:"omitempty"`
	Slug         string `json:"slug" gorm:"column:slug;not null" binding:"omitempty"`
	CoverImage   string `json:"cover_image" gorm:"column:cover_image;" binding:"omitempty"`
	UserId       int    `json:"user_id" gorm:"column:user_id" binding:"omitempty"`
	IsPublish    int    `json:"is_publish" gorm:"column:is_publish" binding:"omitempty"`
	UpdateStatus int    `json:"update_status" gorm:"column:update_status" binding:"omitempty"`
}

func (c *CourseModel) TableName() string {
	return "courses"
}

type CourseList struct {
	Lock  *sync.Mutex
	IdMap map[uint64]*CourseInfo
}

type CourseInfo struct {
	Id           uint64 `json:"id"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	Description  string `json:"description"`
	Slug         string `json:"slug"`
	CoverImage   string `json:"cover_image"`
	UserId       int    `json:"user_id"`
	IsPublish    int    `json:"is_publish"`
	UpdateStatus int    `json:"update_status"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	DeletedAt    string `json:"deleted_at"`
}
