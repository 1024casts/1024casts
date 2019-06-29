package model

import (
	"sync"
	"time"
)

type VideoModel struct {
	BaseModel
	CourseID       uint64    `form:"course_id" gorm:"column:course_id" json:"course_id"`
	SectionID      uint64    `form:"section_id" gorm:"column:section_id" json:"section_id"`
	EpisodeID      int       `form:"episode_id" gorm:"column:episode_id" json:"episode_id"`
	Name           string    `form:"name" gorm:"column:name" json:"name"`
	Description    string    `form:"description" gorm:"column:description" json:"description"`
	CoverImage     string    `form:"cover_image" gorm:"column:cover_image" json:"cover_image"`
	Duration       int       `form:"duration" gorm:"column:duration" json:"duration"`
	DurationStr    string    `form:"duration" gorm:"column:duration" json:"duration"`
	IsFree         int       `form:"is_free" gorm:"column:is_free" json:"is_free"`
	IsPublish      int       `form:"is_publish" gorm:"column:is_publish" json:"is_publish"`
	Mp4URL         string    `form:"mp4_url" gorm:"column:mp4_url" json:"mp4_url"`
	PublishedAt    time.Time `form:"published_at" gorm:"column:published_at" json:"published_at"`
	PublishedAtStr string    `form:"published_at" gorm:"-" json:"-"`
	UserID         int       `form:"user_id" gorm:"column:user_id" json:"user_id"`
}

// TableName sets the insert table name for this struct type
func (v *VideoModel) TableName() string {
	return "videos"
}

type VideoList struct {
	Lock  *sync.Mutex
	IdMap map[uint64]*VideoModel
}
