package repository

import (
	"fmt"

	"github.com/1024casts/1024casts/model"
	"github.com/1024casts/1024casts/pkg/constvar"
)

type VideoRepo struct {
	db *model.Database
}

func NewVideoRepo() *VideoRepo {
	return &VideoRepo{
		db: model.DB,
	}
}

func (repo *VideoRepo) CreateVideo(video model.VideoModel) (id uint64, err error) {
	err = repo.db.Self.Create(&video).Error
	if err != nil {
		return 0, err
	}

	return video.Id, nil
}

func (repo *VideoRepo) GetVideoById(id int) (*model.VideoModel, error) {
	video := model.VideoModel{}
	result := repo.db.Self.Where("id = ?", id).First(&video)

	return &video, result.Error
}

func (repo *VideoRepo) GetVideoByCourseIdAndEpisodeId(courseId uint64, episodeId int) (*model.VideoModel, error) {

	video := new(model.VideoModel)

	if err := repo.db.Self.Where("course_id=?", courseId).Where("episode_id=?", episodeId).Order("id asc").First(&video).Error; err != nil {
		return video, err
	}

	return video, nil
}

func (repo *VideoRepo) GetVideoList(courseId uint64, isGroup bool) ([]*model.VideoModel, error) {

	videos := make([]*model.VideoModel, 0)

	if isGroup {
		if err := repo.db.Self.Where("course_id=?", courseId).Order("id asc").Find(&videos).Error; err != nil {
			return videos, err
		}
	} else {
		if err := repo.db.Self.Where("course_id=? and section_id=0", courseId).Order("id asc").Find(&videos).Error; err != nil {
			return videos, err
		}
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

func (repo *VideoRepo) GetVideoTotalCount() (int, error) {
	var count int
	if err := repo.db.Self.Model(&model.VideoModel{}).Where("is_publish=1").Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (repo *VideoRepo) GetVideoTotalDuration() (int, error) {
	type Result struct {
		Total int
	}
	var result Result
	if err := repo.db.Self.Model(&model.VideoModel{}).Select("sum(duration) as total").Where("is_publish=1").Scan(&result).Error; err != nil {
		return 0, err
	}

	return result.Total, nil
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
