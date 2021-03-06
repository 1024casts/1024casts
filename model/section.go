package model

import "sync"

// User represents a registered user.
type SectionModel struct {
	BaseModel
	Name       string        `json:"name" gorm:"column:name;not null" binding:"required" validate:"min=1,max=32"`
	CourseId   uint64        `json:"course_id" gorm:"column:course_id" binding:"omitempty"`
	Weight     int           `json:"weight" gorm:"column:weight" binding:"omitempty"`
	VideoItems []*VideoModel `json:"video_items" gorm:"-"`
}

func (c *SectionModel) TableName() string {
	return "sections"
}

type SectionList struct {
	Lock  *sync.Mutex
	IdMap map[uint64]*SectionModel
}
