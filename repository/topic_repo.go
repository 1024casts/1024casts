package repository

import (
	"time"

	"github.com/jinzhu/gorm"

	"github.com/1024casts/1024casts/model"
	"github.com/1024casts/1024casts/pkg/constvar"
)

type TopicRepo struct {
	Db *model.Database
}

func NewTopicRepo() *TopicRepo {
	return &TopicRepo{
		Db: model.DB,
	}
}

func (repo *TopicRepo) GetDb() *model.Database {
	return repo.Db
}

func (repo *TopicRepo) CreateTopic(Topic model.TopicModel) (id uint64, err error) {
	err = repo.Db.Self.Create(&Topic).Error
	if err != nil {
		return 0, err
	}

	return Topic.Id, nil
}

func (repo *TopicRepo) AddReply(reply model.ReplyModel) (id uint64, err error) {
	err = repo.Db.Self.Create(&reply).Error
	if err != nil {
		return 0, err
	}

	return reply.Id, nil
}

func (repo *TopicRepo) GetReplyById(id int) (*model.ReplyModel, error) {
	reply := model.ReplyModel{}
	result := repo.Db.Self.Where("id = ?", id).First(&reply)

	return &reply, result.Error
}

func (repo *TopicRepo) GetTopicById(id uint64) (*model.TopicModel, error) {
	Topic := model.TopicModel{}
	result := repo.Db.Self.Where("id = ?", id).First(&Topic)

	return &Topic, result.Error
}

func (repo *TopicRepo) GetTopicList(TopicMap map[string]interface{}, offset, limit int) ([]*model.TopicModel, int, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	Topics := make([]*model.TopicModel, 0)
	var count int

	if err := repo.Db.Self.Model(&model.TopicModel{}).Where(TopicMap).Count(&count).Error; err != nil {
		return Topics, count, err
	}

	if err := repo.Db.Self.Where(TopicMap).Offset(offset).Limit(limit).Order("id desc").Find(&Topics).Error; err != nil {
		return Topics, count, err
	}

	return Topics, count, nil
}

func (repo *TopicRepo) GetTopTopicList(limit int) ([]*model.TopicModel, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	Topics := make([]*model.TopicModel, 0)

	if err := repo.Db.Self.Limit(limit).Order("view_count desc").Find(&Topics).Error; err != nil {
		return Topics, err
	}

	return Topics, nil
}

func (repo *TopicRepo) GetReplyList(replyMap map[string]interface{}, offset, limit int) ([]*model.ReplyModel, int, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	replies := make([]*model.ReplyModel, 0)
	var count int

	if err := repo.Db.Self.Model(&model.ReplyModel{}).Where(replyMap).Count(&count).Error; err != nil {
		return replies, count, err
	}

	if err := repo.Db.Self.Where(replyMap).Offset(offset).Limit(limit).Order("id asc").Find(&replies).Error; err != nil {
		return replies, count, err
	}

	return replies, count, nil
}

func (repo *TopicRepo) GetCategoryList() ([]*model.CategoryModel, error) {
	categories := make([]*model.CategoryModel, 0)

	if err := repo.Db.Self.Order("weight desc").Find(&categories).Error; err != nil {
		return categories, err
	}

	return categories, nil
}

func (repo *TopicRepo) IncrTopicViewCount(topicId uint64) error {
	topic := model.TopicModel{}
	return repo.Db.Self.Model(&topic).Where("id=?", topicId).
		UpdateColumn("view_count", gorm.Expr("view_count + ?", 1)).Error
}

func (repo *TopicRepo) IncrTopicReplyCount(topicId uint64) error {
	topic := model.TopicModel{}
	return repo.Db.Self.Model(&topic).Where("id=?", topicId).
		UpdateColumn("reply_count", gorm.Expr("reply_count + ?", 1)).Error
}

func (repo *TopicRepo) UpdateTopicLastReplyUserId(topicId uint64, userId uint64) error {
	topic := model.TopicModel{}
	topicMap := make(map[string]interface{})
	topicMap["last_reply_user_id"] = userId
	topicMap["last_reply_time_at"] = time.Now()

	return repo.Db.Self.Model(&topic).Where("id=?", topicId).Updates(topicMap).Error
}

func (repo *TopicRepo) IncrReplyLikeCount(replyId int) error {
	reply := model.ReplyModel{}
	return repo.Db.Self.Model(&reply).Where("id=?", replyId).
		UpdateColumn("like_count", gorm.Expr("like_count + ?", 1)).Error
}

func (repo *TopicRepo) UpdateTopic(topicModel model.TopicModel, id uint64) error {
	Topic, err := repo.GetTopicById(id)
	if err != nil {
		return err
	}

	return repo.Db.Self.Model(Topic).Updates(topicModel).Error
}

func (repo *TopicRepo) DeleteTopic(id uint64) error {
	Topic, err := repo.GetTopicById(id)
	if err != nil {
		return err
	}

	return repo.Db.Self.Delete(&Topic).Error
}

func (repo *TopicRepo) Store(Topic *model.TopicModel) (id uint64, err error) {
	//users := model.TopicModel{}

	return 0, nil
}
