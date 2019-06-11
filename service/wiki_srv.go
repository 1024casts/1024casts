package service

import (
	"sync"

	"github.com/1024casts/1024casts/model"
	"github.com/1024casts/1024casts/repository"

	"github.com/lexkong/log"
)

type WikiService struct {
	repo *repository.WikiRepo
}

func NewWikiService() *WikiService {
	return &WikiService{
		repository.NewWikiRepo(),
	}
}

func (srv *WikiService) GetWikiById(id int) (*model.WikiModel, error) {
	Wiki, err := srv.repo.GetWikiById(id)

	if err != nil {
		return Wiki, err
	}

	return Wiki, nil
}

func (srv *WikiService) GetWikiBySlug(slug string) (*model.WikiModel, error) {
	Wiki, err := srv.repo.GetWikiBySlug(slug)

	if err != nil {
		return Wiki, err
	}

	return Wiki, nil
}

func (srv *WikiService) GetWikiList(courseId uint64) ([]*model.WikiModel, error) {
	videos := make([]*model.WikiModel, 0)

	videos, err := srv.repo.GetWikiList(courseId)
	if err != nil {
		log.Warnf("[video] get video list err, course_id: %d", courseId)
		return nil, err
	}

	return videos, nil
}

func (srv *WikiService) GetWikiListPagination(courseId uint64, name string, offset, limit int) ([]*model.WikiModel, uint64, error) {
	infos := make([]*model.WikiModel, 0)

	Wikis, count, err := srv.repo.GetWikiListPagination(courseId, name, offset, limit)
	if err != nil {
		return nil, count, err
	}

	ids := []uint64{}
	for _, Wiki := range Wikis {
		ids = append(ids, Wiki.Id)
	}

	wg := sync.WaitGroup{}
	WikiList := model.WikiList{
		Lock:  new(sync.Mutex),
		IdMap: make(map[uint64]*model.WikiModel, len(Wikis)),
	}

	errChan := make(chan error, 1)
	finished := make(chan bool, 1)

	// Improve query efficiency in parallel
	for _, c := range Wikis {
		wg.Add(1)
		go func(Wiki *model.WikiModel) {
			defer wg.Done()

			//shortId, err := util.GenShortId()
			//if err != nil {
			//	errChan <- err
			//	return
			//}

			WikiList.Lock.Lock()
			defer WikiList.Lock.Unlock()

			WikiList.IdMap[Wiki.Id] = Wiki
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
		infos = append(infos, WikiList.IdMap[id])
	}

	return infos, count, nil
}

func (srv *WikiService) UpdateWiki(WikiMap map[string]interface{}, id int) error {
	err := srv.repo.UpdateWiki(WikiMap, id)

	if err != nil {
		return err
	}

	return nil
}
