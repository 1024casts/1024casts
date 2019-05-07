package model

import (
	"html/template"
	"sync"
	"time"
)

type ReplyModel struct {
	Id         uint64    `gorm:"column:id" json:"id"`
	Body       string    `gorm:"column:body" json:"body"`
	UserId     uint64    `gorm:"column:user_id" json:"user_id"`
	IsBlocked  string    `gorm:"column:is_blocked" json:"is_blocked"`
	OriginBody string    `gorm:"column:origin_body" json:"origin_body"`
	Source     string    `gorm:"column:source" json:"source"`
	TopicID    int       `gorm:"column:topic_id" json:"topic_id"`
	CreatedAt  time.Time `gorm:"column:created_at" json:"created_at"`
	DeletedAt  time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at" json:"updated_at"`
	VoteCount  int       `gorm:"column:vote_count" json:"vote_count"`
}

// TableName sets the insert table name for this struct type
func (r *ReplyModel) TableName() string {
	return "forum_replies"
}

type ReplyList struct {
	Lock  *sync.Mutex
	IdMap map[uint64]*ReplyInfo
}

type ReplyInfo struct {
	Id            uint64        `json:"id"`
	TopicId       int           `json:"topic_id"`
	Body          template.HTML `json:"body"`
	IsBlocked     string        `json:"is_blocked"`
	OriginBody    string        `json:"origin_body"`
	UserID        uint64        `json:"user_id"`
	ReplyUserInfo *UserInfo     `json:"reply_user_info"`
	VoteCount     int           `json:"vote_count"`
	Source        string        `json:"source"`
	CreatedAt     string        `json:"created_at"`
	DeletedAt     string        `json:"deleted_at"`
	UpdatedAt     string        `json:"updated_at"`
}
