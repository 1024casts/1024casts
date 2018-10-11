package model

import (
	"sync"
)

// User represents a registered user.
type CommentModel struct {
	BaseModel
	Type          string `json:"type" gorm:"column:type;not null"`
	RelationId    int    `json:"relation_id" gorm:"column:relation_id" binding:"omitempty"`
	Ip            string `json:"ip" gorm:"column:ip;" binding:"omitempty"`
	Content       string `json:"content" gorm:"column:content;" binding:"omitempty"`
	OriginContent string `json:"origin_content" gorm:"column:origin_content;" binding:"omitempty"`
	UserId        int    `json:"user_id" gorm:"column:user_id" binding:"omitempty"`
	DeviceType    int    `json:"device_type" gorm:"column:device_type" binding:"omitempty"`
}

func (c *CommentModel) TableName() string {
	return "comments"
}

type CommentList struct {
	Lock  *sync.Mutex
	IdMap map[uint64]*CommentModel
}
