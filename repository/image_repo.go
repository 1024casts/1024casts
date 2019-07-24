package repository

import (
	"github.com/1024casts/1024casts/model"
)

type ImageRepo struct {
	db *model.Database
}

func NewImageRepo() *ImageRepo {
	return &ImageRepo{
		db: model.DB,
	}
}

func (repo *ImageRepo) CreateImage(image model.ImageModel) (id uint64, err error) {
	err = repo.db.Self.Create(&image).Error
	if err != nil {
		return 0, err
	}

	return image.ID, nil
}

func (repo *ImageRepo) GetImageById(id int) (*model.ImageModel, error) {
	image := model.ImageModel{}
	result := repo.db.Self.Where("id = ?", id).First(&image)

	return &image, result.Error
}
