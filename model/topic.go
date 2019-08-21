package model

import (
	"html/template"
	"sync"
	"time"
)

type TopicModel struct {
	BaseModel
	Body            string    `gorm:"column:body" json:"body"`
	CategoryID      int       `gorm:"column:category_id" json:"category_id"`
	IsBlocked       string    `gorm:"column:is_blocked" json:"is_blocked"`
	IsExcellent     string    `gorm:"column:is_excellent" json:"is_excellent"`
	LastReplyTimeAt time.Time `gorm:"column:last_reply_time_at" json:"last_reply_time_at"`
	LastReplyUserID uint64    `gorm:"column:last_reply_user_id" json:"last_reply_user_id"`
	OriginBody      string    `gorm:"column:origin_body" json:"origin_body"`
	ReplyCount      int       `gorm:"column:reply_count" json:"reply_count"`
	Source          string    `gorm:"column:source" json:"source"`
	Title           string    `gorm:"column:title" json:"title"`
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
	Id              string         `json:"id"`
	Category        *CategoryModel `json:"category"`
	Title           string         `json:"title"`
	Body            template.HTML  `json:"body"`
	OriginBody      string         `json:"origin_body"`
	Source          string         `json:"source"`
	IsBlocked       string         `json:"is_blocked"`
	IsExcellent     string         `json:"is_excellent"`
	LastReplyTimeAt string         `json:"last_reply_time_at"`
	LastReplyUserId uint64         `json:"last_reply_user_id"`
	LastReplyUser   *UserInfo      `json:"last_reply_user"`
	User            *UserInfo      `json:"user"`
	ViewCount       int            `json:"view_count"`
	VoteCount       int            `json:"vote_count"`
	ReplyCount      int            `json:"reply_count"`
	CreatedAt       string         `json:"created_at"`
	UpdatedAt       string         `json:"updated_at"`
}
