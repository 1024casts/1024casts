package model

import (
	"html/template"
	"sync"
)

// User represents a registered user.
type CommentModel struct {
	BaseModel
	Type          int           `json:"type" gorm:"column:type;not null"`
	RelatedId     int           `json:"related_id" gorm:"column:related_id" binding:"omitempty"`
	Ip            string        `json:"ip" gorm:"column:ip" binding:"omitempty"`
	Content       string        `json:"content" gorm:"column:content" binding:"omitempty"`
	OriginContent string        `json:"origin_content" gorm:"column:origin_content" binding:"omitempty"`
	LikeCount     int           `json:"like_count" gorm:"column:like_count" binding:"omitempty"`
	UserId        uint64        `json:"user_id" gorm:"column:user_id" binding:"omitempty"`
	DeviceType    string        `json:"device_type" gorm:"column:device_type" binding:"omitempty"`
	UserInfo      *UserInfo     `json:"user_info" gorm:"-"`
	ContentHtml   template.HTML `json:"content_html" gorm:"-"`
}

func (c *CommentModel) TableName() string {
	return "comments"
}

type CommentList struct {
	Lock  *sync.Mutex
	IdMap map[uint64]*CommentModel
}
