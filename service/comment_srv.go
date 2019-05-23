package service

import (
	"html/template"
	"sync"

	"github.com/1024casts/1024casts/util"

	"github.com/1024casts/1024casts/model"
	"github.com/1024casts/1024casts/repository"
)

type CommentService struct {
	repo    *repository.CommentRepo
	userSrv *UserService
}

func NewCommentService() *CommentService {
	return &CommentService{
		repository.NewCommentRepo(),
		NewUserService(),
	}
}

func (srv *CommentService) GetCommentById(id int) (*model.CommentModel, error) {
	comment, err := srv.repo.GetCommentById(id)

	if err != nil {
		return comment, err
	}

	return comment, nil
}

func (srv *CommentService) IncrLikeCount(id int) int64 {
	count := srv.repo.IncrLikeCount(id)

	return count
}

func (srv *CommentService) GetCommentList(commentMap map[string]interface{}, offset, limit int) ([]*model.CommentModel, uint64, error) {
	infos := make([]*model.CommentModel, 0)

	comments, count, err := srv.repo.GetCommentList(commentMap, offset, limit)
	if err != nil {
		return nil, count, err
	}

	ids := []uint64{}
	for _, comment := range comments {
		ids = append(ids, comment.Id)
	}

	wg := sync.WaitGroup{}
	commentList := model.CommentList{
		Lock:  new(sync.Mutex),
		IdMap: make(map[uint64]*model.CommentModel, len(comments)),
	}

	errChan := make(chan error, 1)
	finished := make(chan bool, 1)

	// Improve query efficiency in parallel
	for _, c := range comments {
		wg.Add(1)
		go func(comment *model.CommentModel) {
			defer wg.Done()

			//shortId, err := util.GenShortId()
			//if err != nil {
			//	errChan <- err
			//	return
			//}

			commentList.Lock.Lock()
			defer commentList.Lock.Unlock()

			userInfo, _ := srv.userSrv.GetUserById(comment.UserId)

			comment.UserInfo = userInfo
			comment.ContentHtml = template.HTML(comment.Content)
			comment.CreatedAtStr = util.FormatTime(comment.CreatedAt)
			commentList.IdMap[comment.Id] = comment
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
		infos = append(infos, commentList.IdMap[id])
	}

	return infos, count, nil
}

func (srv *CommentService) UpdateComment(commentMap map[string]interface{}, id int) error {
	err := srv.repo.UpdateComment(commentMap, id)

	if err != nil {
		return err
	}

	return nil
}
