package repository

import (
	"time"

	"github.com/1024casts/1024casts/model"
	"github.com/1024casts/1024casts/pkg/constvar"
)

type TopicRepo struct {
	db *model.Database
}

func NewTopicRepo() *TopicRepo {
	return &TopicRepo{
		db: model.DB,
	}
}

func (repo *TopicRepo) GetDb() *model.Database {
	return repo.db
}

func (repo *TopicRepo) CreateTopic(Topic model.TopicModel) (id uint64, err error) {
	err = repo.db.Self.Create(&Topic).Error
	if err != nil {
		return 0, err
	}

	return Topic.Id, nil
}

func (repo *TopicRepo) CreateReply(reply model.ReplyModel) (id uint64, err error) {
	err = repo.db.Self.Create(&reply).Error
	if err != nil {
		return 0, err
	}

	return reply.Id, nil
}

func (repo *TopicRepo) GetReplyById(id int) (*model.ReplyModel, error) {
	reply := model.ReplyModel{}
	result := repo.db.Self.Where("id = ?", id).First(&reply)

	return &reply, result.Error
}

func (repo *TopicRepo) GetTopicById(id uint64) (*model.TopicModel, error) {
	Topic := model.TopicModel{}
	result := repo.db.Self.Where("id = ?", id).First(&Topic)

	return &Topic, result.Error
}

func (repo *TopicRepo) GetTopicList(TopicMap map[string]interface{}, offset, limit int) ([]*model.TopicModel, int, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	Topics := make([]*model.TopicModel, 0)
	var count int

	if err := repo.db.Self.Model(&model.TopicModel{}).Where(TopicMap).Count(&count).Error; err != nil {
		return Topics, count, err
	}

	if err := repo.db.Self.Where(TopicMap).Offset(offset).Limit(limit).Order("id desc").Find(&Topics).Error; err != nil {
		return Topics, count, err
	}

	return Topics, count, nil
}

func (repo *TopicRepo) GetReplyList(replyMap map[string]interface{}, offset, limit int) ([]*model.ReplyModel, int, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	replies := make([]*model.ReplyModel, 0)
	var count int

	if err := repo.db.Self.Model(&model.ReplyModel{}).Where(replyMap).Count(&count).Error; err != nil {
		return replies, count, err
	}

	if err := repo.db.Self.Where(replyMap).Offset(offset).Limit(limit).Order("id desc").Find(&replies).Error; err != nil {
		return replies, count, err
	}

	return replies, count, nil
}

func (repo *TopicRepo) GetCategoryList() ([]*model.CategoryModel, error) {
	categories := make([]*model.CategoryModel, 0)

	if err := repo.db.Self.Order("weight desc").Find(&categories).Error; err != nil {
		return categories, err
	}

	return categories, nil
}

func (repo *TopicRepo) IncrTopicViewCount(topicId uint64) error {
	topic := model.TopicModel{}
	topicMap := make(map[string]interface{})
	topicMap["view_count"] = topic.ViewCount + 1

	return repo.db.Self.Model(&topic).Where("id=?", topicId).Updates(topicMap).Error
}

func (repo *TopicRepo) IncrTopicReplyCount(topicId uint64) error {
	topic := model.TopicModel{}
	topicMap := make(map[string]interface{})
	topicMap["reply_count"] = topic.ReplyCount + 1

	return repo.db.Self.Model(&topic).Where("id=?", topicId).Updates(topicMap).Error
}

func (repo *TopicRepo) UpdateTopicLastReplyUserId(topicId uint64, userId uint64) error {
	topic := model.TopicModel{}
	topicMap := make(map[string]interface{})
	topicMap["last_reply_user_id"] = userId
	topicMap["last_reply_time_at"] = time.Now()

	return repo.db.Self.Model(&topic).Where("id=?", topicId).Updates(topicMap).Error
}

func (repo *TopicRepo) IncrReplyLikeCount(id int) error {
	reply := model.ReplyModel{}
	replyMap := make(map[string]interface{})
	replyMap["like_count"] = reply.LikeCount + 1

	return repo.db.Self.Model(&reply).Where("id=?", id).Updates(replyMap).Error
}

func (repo *TopicRepo) UpdateTopic(topicModel model.TopicModel, id uint64) error {
	Topic, err := repo.GetTopicById(id)
	if err != nil {
		return err
	}

	return repo.db.Self.Model(Topic).Updates(topicModel).Error
}

func (repo *TopicRepo) DeleteTopic(id uint64) error {
	Topic, err := repo.GetTopicById(id)
	if err != nil {
		return err
	}

	return repo.db.Self.Delete(&Topic).Error
}

func (repo *TopicRepo) Store(Topic *model.TopicModel) (id uint64, err error) {
	//users := model.TopicModel{}

	return 0, nil
}
