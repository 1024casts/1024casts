package repository

import (
	"fmt"

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

func (repo *WikiRepo) GetWikiById(id int) (*model.WikiModel, error) {
	video := model.WikiModel{}
	result := repo.db.Self.Where("id = ?", id).First(&video)

	return &video, result.Error
}

func (repo *WikiRepo) GetWikiBySlug(slug string) (*model.WikiModel, error) {
	video := model.WikiModel{}
	result := repo.db.Self.Where("slug = ?", slug).First(&video)

	return &video, result.Error
}

func (repo *WikiRepo) GetWikiList(courseId uint64) ([]*model.WikiModel, error) {

	videos := make([]*model.WikiModel, 0)

	if err := repo.db.Self.Where("course_id=?", courseId).Order("id asc").Find(&videos).Error; err != nil {
		return videos, err
	}

	return videos, nil
}

func (repo *WikiRepo) GetWikiListPagination(courseId uint64, name string, offset, limit int) ([]*model.WikiModel, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	Wikis := make([]*model.WikiModel, 0)
	var count uint64

	where := fmt.Sprintf("name like '%%%s%%'", name)
	if err := repo.db.Self.Model(&model.WikiModel{}).Where("course_id=?", courseId).Where(where).Count(&count).Error; err != nil {
		return Wikis, count, err
	}

	if err := repo.db.Self.Where("course_id=?", courseId).Where(where).Offset(offset).Limit(limit).Order("id desc").Find(&Wikis).Error; err != nil {
		return Wikis, count, err
	}

	return Wikis, count, nil
}

func (repo *WikiRepo) UpdateWiki(WikiMap map[string]interface{}, id int) error {

	Wiki, err := repo.GetWikiById(id)
	if err != nil {
		return err
	}

	return repo.db.Self.Model(Wiki).Updates(WikiMap).Error
}

func (repo *WikiRepo) DeleteWiki(id int) error {
	Wiki, err := repo.GetWikiById(id)
	if err != nil {
		return err
	}

	return repo.db.Self.Delete(&Wiki).Error
}
