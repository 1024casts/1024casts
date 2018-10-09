package model

import "sync"

// User represents a registered user.
type SectionModel struct {
	BaseModel
	Name     string `json:"name" gorm:"column:name;not null" binding:"required" validate:"min=1,max=32"`
	CourseId int    `json:"course_id" gorm:"column:course_id" binding:"omitempty"`
	Order    int    `json:"order" gorm:"column:order" binding:"omitempty"`
}

func (c *SectionModel) TableName() string {
	return "sections"
}

type SectionList struct {
	Lock  *sync.Mutex
	IdMap map[uint64]*SectionModel
}
