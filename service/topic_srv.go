package service

import (
	"sync"

	"github.com/1024casts/1024casts/model"
	"github.com/1024casts/1024casts/repository"
)

type TopicService struct {
	repo     *repository.TopicRepo
	videoSrv *VideoService
}

func NewTopicService() *TopicService {
	return &TopicService{
		repo:     repository.NewTopicRepo(),
		videoSrv: NewVideoService(),
	}
}

func (srv *TopicService) CreateTopic(user model.TopicModel) (id uint64, err error) {
	id, err = srv.repo.CreateTopic(user)

	if err != nil {
		return id, err
	}

	return id, nil
}

func (srv *TopicService) GetTopicById(id int) (*model.TopicModel, error) {
	Topic, err := srv.repo.GetTopicById(id)

	if err != nil {
		return Topic, err
	}

	return Topic, nil
}

func (srv *TopicService) GetTopicList(TopicMap map[string]interface{}, offset, limit int) ([]*model.TopicModel, uint64, error) {
	infos := make([]*model.TopicModel, 0)

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
		IdMap: make(map[uint64]*model.TopicModel, len(Topics)),
	}

	errChan := make(chan error, 1)
	finished := make(chan bool, 1)

	// Improve query efficiency in parallel
	for _, c := range Topics {
		wg.Add(1)
		go func(Topic *model.TopicModel) {
			defer wg.Done()

			//shortId, err := util.GenShortId()
			//if err != nil {
			//	errChan <- err
			//	return
			//}

			TopicList.Lock.Lock()
			defer TopicList.Lock.Unlock()

			TopicList.IdMap[Topic.Id] = Topic
		}(c)
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

func (srv *TopicService) UpdateTopic(TopicMap map[string]interface{}, id int) error {
	err := srv.repo.UpdateTopic(TopicMap, id)

	if err != nil {
		return err
	}

	return nil
}
