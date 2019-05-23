package repository

import (
	"github.com/1024casts/1024casts/model"
	"github.com/1024casts/1024casts/pkg/constvar"
	"github.com/jinzhu/gorm"
)

type CommentRepo struct {
	db *model.Database
}

func NewCommentRepo() *CommentRepo {
	return &CommentRepo{
		db: model.DB,
	}
}

func (repo *CommentRepo) GetCommentById(id int) (*model.CommentModel, error) {
	comment := model.CommentModel{}
	result := repo.db.Self.Where("id = ?", id).First(&comment)

	return &comment, result.Error
}

func (repo *CommentRepo) IncrLikeCount(id int) int64 {
	result := repo.db.Self.Model(&model.CommentModel{}).Where("id = ?", id).
		UpdateColumn("like_count", gorm.Expr("like_count + ?", 1))

	return result.RowsAffected
}

func (repo *CommentRepo) GetCommentList(courseMap map[string]interface{}, offset, limit int) ([]*model.CommentModel, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	comments := make([]*model.CommentModel, 0)
	var count uint64

	if err := repo.db.Self.Model(&model.CommentModel{}).Where(courseMap).Count(&count).Error; err != nil {
		return comments, count, err
	}

	if err := repo.db.Self.Where(courseMap).Offset(offset).Limit(limit).Order("id asc").Find(&comments).Error; err != nil {
		return comments, count, err
	}

	return comments, count, nil
}

func (repo *CommentRepo) UpdateComment(commentMap map[string]interface{}, id int) error {

	comment, err := repo.GetCommentById(id)
	if err != nil {
		return err
	}

	return repo.db.Self.Model(comment).Updates(commentMap).Error
}

func (repo *CommentRepo) DeleteComment(id int) error {
	comment, err := repo.GetCommentById(id)
	if err != nil {
		return err
	}

	return repo.db.Self.Delete(&comment).Error
}
