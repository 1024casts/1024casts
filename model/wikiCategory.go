package model

import "sync"

type WikiCategoryModel struct {
	BaseModel
	Name        string           `gorm:"column:name" json:"name"`
	Description string           `gorm:"column:description" json:"description"`
	Weight      int              `gorm:"column:weight" json:"weight"`
	WikiPages   []*WikiPageModel `json:"wiki_pages" gorm:"-"`
}

// TableName sets the insert table name for this struct type
func (c *WikiCategoryModel) TableName() string {
	return "wiki_categories"
}

type WikiCategoryList struct {
	Lock  *sync.Mutex
	IdMap map[uint64]*WikiCategoryModel
}
