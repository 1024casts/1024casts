package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"

	"github.com/1024casts/1024casts/model"
	"github.com/1024casts/1024casts/pkg/constvar"
)

type WikiRepo struct {
	db *model.Database
}

func NewWikiRepo() *WikiRepo {
	return &WikiRepo{
		db: model.DB,
	}
}

func (repo *WikiRepo) GetCategoryList() ([]*model.WikiCategoryModel, error) {
	categories := make([]*model.WikiCategoryModel, 0)

	if err := repo.db.Self.Where("status=1").Order("weight desc").Find(&categories).Error; err != nil {
		return categories, err
	}

	return categories, nil
}

func (repo *WikiRepo) GetWikiById(id int) (*model.WikiPageModel, error) {
	page := model.WikiPageModel{}
	result := repo.db.Self.Where("id = ?", id).First(&page)

	return &page, result.Error
}

func (repo *WikiRepo) GetWikiBySlug(slug string) (*model.WikiPageModel, error) {
	page := model.WikiPageModel{}
	result := repo.db.Self.Where("slug = ?", slug).First(&page)

	return &page, result.Error
}

func (repo *WikiRepo) GetWikiPageListByCategoryId(categoryId uint64) ([]*model.WikiPageModel, error) {
	pages := make([]*model.WikiPageModel, 0)

	if err := repo.db.Self.Where("is_parent=0 AND category_id=? AND status=1", categoryId).Order("weight desc").Find(&pages).Error; err != nil {
		return pages, err
	}

	return pages, nil
}

func (repo *WikiRepo) GetWikiListPagination(courseId uint64, name string, offset, limit int) ([]*model.WikiPageModel, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	pages := make([]*model.WikiPageModel, 0)
	var count uint64

	where := fmt.Sprintf("name like '%%%s%%'", name)
	if err := repo.db.Self.Model(&model.WikiPageModel{}).Where("course_id=?", courseId).Where(where).Count(&count).Error; err != nil {
		return pages, count, err
	}

	if err := repo.db.Self.Where("course_id=?", courseId).Where(where).Offset(offset).Limit(limit).Order("id desc").Find(&pages).Error; err != nil {
		return pages, count, err
	}

	return pages, count, nil
}

func (repo *WikiRepo) UpdateWiki(WikiMap map[string]interface{}, id int) error {

	page, err := repo.GetWikiById(id)
	if err != nil {
		return err
	}

	return repo.db.Self.Model(page).Updates(WikiMap).Error
}

func (repo *WikiRepo) IncrViewCount(id uint64) error {
	wiki := model.WikiPageModel{}
	return repo.db.Self.Model(&wiki).Where("id=?", id).UpdateColumn("view_count", gorm.Expr("view_count + ?", 1)).Error
}

func (repo *WikiRepo) DeleteWiki(id int) error {
	page, err := repo.GetWikiById(id)
	if err != nil {
		return err
	}

	return repo.db.Self.Delete(&page).Error
}
