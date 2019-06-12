package model

import "sync"

type WikiPageModel struct {
	BaseModel
	CategoryId    uint64 `gorm:"column:category_id" json:"category_id"`
	Slug          string `gorm:"column:slug" json:"slug"`
	Title         string `gorm:"column:title" json:"title"`
	Summary       string `gorm:"column:summary" json:"summary"`
	OriginContent string `gorm:"column:origin_content" json:"origin_content"`
	Content       string `gorm:"column:content" json:"content"`
	ViewCount     int    `gorm:"column:view_count" json:"view_count"`
	CommentCount  int    `gorm:"column:comment_count" json:"comment_count"`
	FixCount      int    `gorm:"column:fix_count" json:"fix_count"`
	IsShow        int    `gorm:"column:is_show" json:"is_show"`
	Status        int    `gorm:"column:status" json:"status"`
	UserId        int    `gorm:"column:user_id" json:"user_id"`
}

// TableName sets the insert table name for this struct type
func (w *WikiPageModel) TableName() string {
	return "wiki_pages"
}

type WikiList struct {
	Lock  *sync.Mutex
	IdMap map[uint64]*WikiPageModel
}
