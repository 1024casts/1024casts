package service

import (
	"sync"
	"time"

	"1024casts/backend/model"
	"1024casts/backend/repository"
)

type UserService struct {
	userRepo *repository.UserRepo
}

func NewUserService() *UserService {
	return &UserService{
		repository.NewUserRepo(),
	}
}

func (srv *UserService) CreateUser(user model.UserModel) (id uint64, err error) {
	id, err = srv.userRepo.CreateUser(user)

	if err != nil {
		return id, err
	}

	return id, nil
}

func (srv *UserService) GetUserById(id uint64) (*model.UserModel, error) {
	user, err := srv.userRepo.GetUserById(id)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (srv *UserService) GetUserByUsername(username string) (*model.UserModel, error) {
	user, err := srv.userRepo.GetUserByUsername(username)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (srv *UserService) GetUserList(username string, offset, limit int) ([]*model.UserModel, uint64, error) {
	infos := make([]*model.UserModel, 0)
	users, count, err := srv.userRepo.GetUserList(username, offset, limit)
	if err != nil {
		return nil, count, err
	}

	ids := []uint64{}
	for _, user := range users {
		ids = append(ids, user.Id)
	}

	wg := sync.WaitGroup{}
	userList := model.UserList{
		Lock:  new(sync.Mutex),
		IdMap: make(map[uint64]*model.UserModel, len(users)),
	}

	errChan := make(chan error, 1)
	finished := make(chan bool, 1)

	// Improve query efficiency in parallel
	for _, u := range users {
		wg.Add(1)
		go func(u *model.UserModel) {
			defer wg.Done()

			userList.Lock.Lock()
			defer userList.Lock.Unlock()

			u.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", u.CreatedAt.String())
			userList.IdMap[u.Id] = u
		}(u)
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
		infos = append(infos, userList.IdMap[id])
	}

	return infos, count, nil
}

func (srv *UserService) UpdateUser(userMap map[string]interface{}, id uint64) error {
	err := srv.userRepo.Update(userMap, id)

	if err != nil {
		return err
	}

	return nil
}

func (srv *UserService) DeleteUser(id uint64) error {
	err := srv.userRepo.DeleteUser(id)

	if err != nil {
		return err
	}

	return nil
}
