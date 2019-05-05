package service

import (
	"html/template"
	"sync"

	"github.com/1024casts/1024casts/util"

	"github.com/1024casts/1024casts/model"
	"github.com/1024casts/1024casts/repository"
)

type TopicService struct {
	repo     *repository.TopicRepo
	videoSrv *VideoService
	userSrv  *UserService
}

func NewTopicService() *TopicService {
	return &TopicService{
		repo:     repository.NewTopicRepo(),
		videoSrv: NewVideoService(),
		userSrv:  NewUserService(),
	}
}

func (srv *TopicService) CreateTopic(topic model.TopicModel) (id uint64, err error) {
	id, err = srv.repo.CreateTopic(topic)

	if err != nil {
		return id, err
	}

	return id, nil
}

func (srv *TopicService) GetTopicById(id int) (*model.TopicInfo, error) {
	topicModel, err := srv.repo.GetTopicById(id)
	topic := srv.trans(topicModel)

	if err != nil {
		return topic, err
	}

	return topic, nil
}

func (srv *TopicService) GetTopicList(TopicMap map[string]interface{}, offset, limit int) ([]*model.TopicInfo, int, error) {
	infos := make([]*model.TopicInfo, 0)

	Topics, count, err := srv.repo.GetTopicList(TopicMap, offset, limit)
	if err != nil {
		return nil, count, err
	}

	ids := []uint64{}
	for _, Topic := range Topics {
		ids = append(ids, Topic.Id)
	}

	wg := sync.WaitGroup{}
	TopicList := model.TopicList{
		Lock:  new(sync.Mutex),
		IdMap: make(map[uint64]*model.TopicInfo, len(Topics)),
	}

	errChan := make(chan error, 1)
	finished := make(chan bool, 1)

	// Improve query efficiency in parallel
	for _, t := range Topics {
		wg.Add(1)
		go func(Topic *model.TopicModel) {
			defer wg.Done()

			TopicList.Lock.Lock()
			defer TopicList.Lock.Unlock()

			TopicList.IdMap[Topic.Id] = srv.trans(Topic)
		}(t)
	}

	go func() {
		wg.Wait()
		close(finished)
	}()

	select {
	case <-finished:
	case err := <-errChan:
		return nil, count, err
	}

	for _, id := range ids {
		infos = append(infos, TopicList.IdMap[id])
	}

	return infos, count, nil
}

func (srv *TopicService) trans(topic *model.TopicModel) *model.TopicInfo {
	lastReplyUser, _ := srv.userSrv.GetUserById(topic.LastReplyUserID)
	creator, _ := srv.userSrv.GetUserById(topic.UserID)
	return &model.TopicInfo{
		Id:                topic.Id,
		CategoryID:        topic.CategoryID,
		Title:             topic.Title,
		Body:              template.HTML(topic.Body),
		OriginBody:        topic.OriginBody,
		Source:            topic.Source,
		IsBlocked:         topic.IsBlocked,
		IsExcellent:       topic.IsExcellent,
		LastReplyTimeAt:   util.FormatTime(topic.LastReplyTimeAt),
		LastReplyUserId:   topic.LastReplyUserID,
		LastReplyUserInfo: lastReplyUser,
		UserInfo:          creator,
		ViewCount:         topic.ViewCount,
		VoteCount:         topic.VoteCount,
		ReplyCount:        topic.ReplyCount,
		CreatedAt:         util.TimeToDateString(topic.CreatedAt),
		UpdatedAt:         util.TimeToString(topic.UpdatedAt),
	}
}

func (srv *TopicService) transReply(reply *model.ReplyModel) *model.ReplyInfo {
	replyUser, _ := srv.userSrv.GetUserById(reply.UserId)
	return &model.ReplyInfo{
		Id:            reply.Id,
		TopicId:       reply.TopicID,
		Body:          template.HTML(reply.Body),
		IsBlocked:     reply.IsBlocked,
		OriginBody:    reply.OriginBody,
		UserID:        reply.UserId,
		ReplyUserInfo: replyUser,
		VoteCount:     reply.VoteCount,
		Source:        reply.Source,
		CreatedAt:     util.FormatTime(reply.CreatedAt),
		DeletedAt:     "",
		UpdatedAt:     util.TimeToString(reply.UpdatedAt),
	}
}

func (srv *TopicService) UpdateTopic(TopicMap map[string]interface{}, id int) error {
	err := srv.repo.UpdateTopic(TopicMap, id)

	if err != nil {
		return err
	}

	return nil
}

func (srv *TopicService) IncrTopicViewCount(id int) error {
	err := srv.repo.IncrTopicViewCount(id)

	if err != nil {
		return err
	}

	return nil
}

func (srv *TopicService) GetReplyList(replyMap map[string]interface{}, offset, limit int) ([]*model.ReplyInfo, int, error) {
	infos := make([]*model.ReplyInfo, 0)

	replies, count, err := srv.repo.GetReplyList(replyMap, offset, limit)
	if err != nil {
		return nil, count, err
	}

	ids := []uint64{}
	for _, reply := range replies {
		ids = append(ids, reply.Id)
	}

	wg := sync.WaitGroup{}
	replyList := model.ReplyList{
		Lock:  new(sync.Mutex),
		IdMap: make(map[uint64]*model.ReplyInfo, len(replies)),
	}

	errChan := make(chan error, 1)
	finished := make(chan bool, 1)

	// Improve query efficiency in parallel
	for _, r := range replies {
		wg.Add(1)
		go func(reply *model.ReplyModel) {
			defer wg.Done()

			replyList.Lock.Lock()
			defer replyList.Lock.Unlock()

			replyList.IdMap[reply.Id] = srv.transReply(reply)
		}(r)
	}

	go func() {
		wg.Wait()
		close(finished)
	}()

	select {
	case <-finished:
	case err := <-errChan:
		return nil, count, err
	}

	for _, id := range ids {
		infos = append(infos, replyList.IdMap[id])
	}

	return infos, count, nil
}

func (srv *TopicService) GetCategoryList() ([]*model.CategoryModel, error) {
	categories, err := srv.repo.GetCategoryList()
	if err != nil {
		return nil, err
	}

	return categories, nil
}
