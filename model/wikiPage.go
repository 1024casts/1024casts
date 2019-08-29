package model

import (
	"html/template"
	"sync"
)

type WikiPageModel struct {
	BaseModel
	CategoryId    uint64 `gorm:"column:category_id" json:"category_id"`
	IsParent      int    `gorm:"column:is_parent" json:"is_parent"`
	Slug          string `gorm:"column:slug" json:"slug"`
	Title         string `gorm:"column:title" json:"title"`
	Summary       string `gorm:"column:summary" json:"summary"`
	OriginContent string `gorm:"column:origin_content" json:"origin_content"`
	Content       string `gorm:"column:content" json:"content"`
	Weight        int    `gorm:"column:weight" json:"weight"`
	ViewCount     int    `gorm:"column:view_count" json:"view_count"`
	CommentCount  int    `gorm:"column:comment_count" json:"comment_count"`
	FixCount      int    `gorm:"column:fix_count" json:"fix_count"`
	Status        int    `gorm:"column:status" json:"status"`
	UserId        uint64 `gorm:"column:user_id" json:"user_id"`
}

// TableName sets the insert table name for this struct type
func (w *WikiPageModel) TableName() string {
	return "wiki_pages"
}

type WikiList struct {
	Lock  *sync.Mutex
	IdMap map[uint64]*WikiPageModel
}

type WikiPageInfo struct {
	Id            uint64        `json:"id"`
	CategoryId    uint64        `json:"category_id"`
	IsParent      int           `json:"is_parent"`
	Slug          string        `json:"slug"`
	Title         string        `json:"title"`
	Summary       string        `json:"summary"`
	OriginContent string        `json:"origin_content"`
	Content       template.HTML `json:"content"`
	Weight        int           `json:"weight"`
	ViewCount     int           `json:"view_count"`
	CommentCount  int           `json:"comment_count"`
	FixCount      int           `json:"fix_count"`
	Status        int           `json:"status"`
	UserId        uint64        `json:"user_id"`
	CreatedAt     string        `json:"created_at"`
	UpdatedAt     string        `json:"updated_at"`
}
