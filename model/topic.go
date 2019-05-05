package model

import (
	"html/template"
	"sync"
	"time"
)

type TopicModel struct {
	Body            string    `gorm:"column:body" json:"body"`
	CategoryID      int       `gorm:"column:category_id" json:"category_id"`
	CreatedAt       time.Time `gorm:"column:created_at" json:"created_at"`
	DeletedAt       time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	Id              uint64    `gorm:"column:id" json:"id"`
	IsBlocked       string    `gorm:"column:is_blocked" json:"is_blocked"`
	IsExcellent     string    `gorm:"column:is_excellent" json:"is_excellent"`
	LastReplyTimeAt time.Time `gorm:"column:last_reply_time_at" json:"last_reply_time_at"`
	LastReplyUserID uint64    `gorm:"column:last_reply_user_id" json:"last_reply_user_id"`
	OriginBody      string    `gorm:"column:origin_body" json:"origin_body"`
	ReplyCount      int       `gorm:"column:reply_count" json:"reply_count"`
	Source          string    `gorm:"column:source" json:"source"`
	Title           string    `gorm:"column:title" json:"title"`
	UpdatedAt       time.Time `gorm:"column:updated_at" json:"updated_at"`
	UserID          uint64    `gorm:"column:user_id" json:"user_id"`
	ViewCount       int       `gorm:"column:view_count" json:"view_count"`
	VoteCount       int       `gorm:"column:vote_count" json:"vote_count"`
}

// TableName sets the insert table name for this struct type
func (t *TopicModel) TableName() string {
	return "forum_topics"
}

type TopicList struct {
	Lock  *sync.Mutex
	IdMap map[uint64]*TopicInfo
}

type TopicInfo struct {
	Id                uint64        `json:"id"`
	CategoryID        int           `json:"category_id"`
	Title             string        `json:"title"`
	Body              template.HTML `json:"body"`
	OriginBody        string        `json:"origin_body"`
	Source            string        `json:"source"`
	IsBlocked         string        `json:"is_blocked"`
	IsExcellent       string        `json:"is_excellent"`
	LastReplyTimeAt   string        `json:"last_reply_time_at"`
	LastReplyUserId   uint64        `json:"last_reply_user_id"`
	LastReplyUserInfo *UserInfo     `json:"last_reply_user_info"`
	UserInfo          *UserInfo     `json:"user_info"`
	ViewCount         int           `json:"view_count"`
	VoteCount         int           `json:"vote_count"`
	ReplyCount        int           `json:"reply_count"`
	CreatedAt         string        `json:"created_at"`
	UpdatedAt         string        `json:"updated_at"`
}
