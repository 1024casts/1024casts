package service

import (
	"html/template"
	"sync"

	"github.com/lexkong/log"

	"github.com/1024casts/1024casts/util"

	"github.com/1024casts/1024casts/model"
	"github.com/1024casts/1024casts/repository"
)

type TopicService struct {
	repo     *repository.TopicRepo
	videoSrv *VideoService
	userSrv  *userService
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

func (srv *TopicService) AddReply(userId, topicId uint64, originBody string) (id uint64, err error) {
	// todo: 使用事务进行批量处理

	replyModel := model.ReplyModel{
		TopicId:    topicId,
		OriginBody: originBody,
		Body:       util.MarkdownToHtml(util.ParseMentionUser(originBody)),
		Source:     "PC",
		IsBlocked:  "no",
		UserId:     userId,
	}

	// create reply
	id, err = srv.repo.AddReply(replyModel)
	if err != nil {
		log.Warnf("[topic] create reply err: %+v", err)
		return id, err
	}

	// update reply count
	err = srv.IncrTopicReplyCount(topicId)
	if err != nil {
		log.Warnf("[topic] incr reply count err: %+v", err)
		return id, err
	}

	// update last_reply_uid
	err = srv.UpdateTopicLastReplyUserId(topicId, userId)
	if err != nil {
		log.Warnf("[topic] incr last reply user id err: %+v", err)
		return id, err
	}

	// update user
	err = srv.userSrv.IncrReplyCount(userId)
	if err != nil {
		log.Warnf("[topic] incr user reply count err: %+v", err)
		return id, err
	}

	// 发送通知

	return id, nil
}

func (srv *TopicService) GetTopicById(id uint64) (*model.TopicInfo, error) {
	topicModel, err := srv.repo.GetTopicById(id)
	topic := srv.trans(topicModel)

	if err != nil {
		return topic, err
	}

	return topic, nil
}

// 获取热门topic
func (srv *TopicService) GetTopTopicList(limit int) ([]*model.TopicInfo, error) {
	infos := make([]*model.TopicInfo, 0)

	Topics, err := srv.repo.GetTopTopicList(limit)
	if err != nil {
		return nil, err
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
		return nil, err
	}

	for _, id := range ids {
		infos = append(infos, TopicList.IdMap[id])
	}

	return infos, nil
}

func (srv *TopicService) GetTopicList(topicMap map[string]interface{}, offset, limit int) ([]*model.TopicInfo, int, error) {
	infos := make([]*model.TopicInfo, 0)

	Topics, count, err := srv.repo.GetTopicList(topicMap, offset, limit)
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
		Id:                util.EncodeTopicId(int64(topic.Id)),
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
		TopicId:       reply.TopicId,
		Body:          template.HTML(reply.Body),
		IsBlocked:     reply.IsBlocked,
		OriginBody:    reply.OriginBody,
		UserID:        reply.UserId,
		ReplyUserInfo: replyUser,
		LikeCount:     reply.LikeCount,
		Source:        reply.Source,
		CreatedAt:     util.FormatTime(reply.CreatedAt),
		DeletedAt:     "",
		UpdatedAt:     util.TimeToString(reply.UpdatedAt),
	}
}

func (srv *TopicService) UpdateTopic(topicModel model.TopicModel, id uint64) error {
	err := srv.repo.UpdateTopic(topicModel, id)

	if err != nil {
		return err
	}

	return nil
}

func (srv *TopicService) IncrTopicViewCount(id uint64) error {
	err := srv.repo.IncrTopicViewCount(id)
	if err != nil {
		return err
	}
	return nil
}

func (srv *TopicService) IncrTopicReplyCount(id uint64) error {
	err := srv.repo.IncrTopicReplyCount(id)
	if err != nil {
		return err
	}
	return nil
}

func (srv *TopicService) UpdateTopicLastReplyUserId(id uint64, userId uint64) error {
	err := srv.repo.UpdateTopicLastReplyUserId(id, userId)
	if err != nil {
		return err
	}
	return nil
}

func (srv *TopicService) IncrReplyLikeCount(id int) error {
	err := srv.repo.IncrReplyLikeCount(id)
	if err != nil {
		return err
	}
	return nil
}

func (srv *TopicService) GetReplyById(id int) (*model.ReplyInfo, error) {
	replyModel, err := srv.repo.GetReplyById(id)
	reply := srv.transReply(replyModel)

	if err != nil {
		return reply, err
	}
	return reply, nil
}

func (srv *TopicService) GetReplyList(replyMap map[string]interface{}, offset, limit int) ([]*model.ReplyInfo, int, error) {
	infos := make([]*model.ReplyInfo, 0)

	replies, count, err := srv.repo.GetReplyList(replyMap, offset, limit)
	if err != nil {
		return nil, count, err
	}

	var ids []uint64
	for _, reply := range replies {
		ids = append(ids, reply.Id)
	}

	var wg sync.WaitGroup
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
