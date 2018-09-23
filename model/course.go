package model

import (
	"1024casts/backend/pkg/constvar"
	"sync"
)

// User represents a registered user.
type CourseModel struct {
	BaseModel
	Name        string `json:"name" gorm:"column:name;not null" binding:"required" validate:"min=1,max=32"`
	Type        string `json:"type" gorm:"column:type;not null"`
	Description string `json:"description" gorm:"column:description;" binding:"omitempty"`
	Slug        string `json:"slug" gorm:"column:slug;not null" binding:"omitempty"`
	CoverImage  string `json:"cover_image" gorm:"column:cover_image;" binding:"omitempty"`
	UserId      int    `json:"user_id" gorm:"column:user_id" binding:"omitempty"`
	IsPublish   int    `json:"is_publish" gorm:"column:is_publish" binding:"omitempty"`
}

func (c *CourseModel) TableName() string {
	return "courses"
}

type CourseList struct {
	Lock  *sync.Mutex
	IdMap map[uint64]*CourseModel
}

// ListUser List all users
func ListCourse(course CourseModel, offset, limit int) ([]*CourseModel, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	courses := make([]*CourseModel, 0)
	var count uint64

	where :=  "is_publish >=0 "
	if err := DB.Self.Model(&CourseModel{}).Where(where).Count(&count).Error; err != nil {
		return courses, count, err
	}

	if err := DB.Self.Where(where).Offset(offset).Limit(limit).Order("id desc").Find(&courses).Error; err != nil {
		return courses, count, err
	}

	return courses, count, nil
}
