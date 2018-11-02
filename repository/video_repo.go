package repository

import (
	"1024casts/backend/model"
	"1024casts/backend/pkg/constvar"
	"fmt"
)

type VideoRepo struct {
	db *model.Database
}

func NewVideoRepo() *VideoRepo {
	return &VideoRepo{
		db: model.DB,
	}
}

func (repo *VideoRepo) GetVideoById(id int) (*model.VideoModel, error) {
	Video := model.VideoModel{}
	result := repo.db.Self.Where("id = ?", id).First(&Video)

	return &Video, result.Error
}

func (repo *VideoRepo) GetVideoList(courseId uint64) ([]*model.VideoModel, error) {

	videos := make([]*model.VideoModel, 0)

	if err := repo.db.Self.Where("course_id=?", courseId).Where("course_id=?", courseId).Order("id asc").Find(&videos).Error; err != nil {
		return videos, err
	}

	return videos, nil
}

func (repo *VideoRepo) GetVideoListPagination(courseId uint64, name string, offset, limit int) ([]*model.VideoModel, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	Videos := make([]*model.VideoModel, 0)
	var count uint64

	where := fmt.Sprintf("name like '%%%s%%'", name)
	if err := repo.db.Self.Model(&model.VideoModel{}).Where("course_id=?", courseId).Where(where).Count(&count).Error; err != nil {
		return Videos, count, err
	}

	if err := repo.db.Self.Where("course_id=?", courseId).Where(where).Offset(offset).Limit(limit).Order("id desc").Find(&Videos).Error; err != nil {
		return Videos, count, err
	}

	return Videos, count, nil
}

func (repo *VideoRepo) UpdateVideo(VideoMap map[string]interface{}, id int) error {

	Video, err := repo.GetVideoById(id)
	if err != nil {
		return err
	}

	return repo.db.Self.Model(Video).Updates(VideoMap).Error
}

func (repo *VideoRepo) DeleteVideo(id int) error {
	Video, err := repo.GetVideoById(id)
	if err != nil {
		return err
	}

	return repo.db.Self.Delete(&Video).Error
}
