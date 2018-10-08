package service

import (
	"sync"

	"1024casts/backend/model"
	"1024casts/backend/repository"
)

type CourseService struct {
	courseRepo *repository.CourseRepo
}

func NewCourseService() *CourseService {
	return &CourseService{
		repository.NewCourseRepo(),
	}
}

func (srv *CourseService) CreateCourse(user model.CourseModel) (id uint64, err error) {
	id, err = srv.courseRepo.CreateCourse(user)

	if err != nil {
		return id, err
	}

	return id, nil
}

func (srv *CourseService) GetCourseById(id int) (*model.CourseModel, error) {
	course, err := srv.courseRepo.GetCourseById(id)

	if err != nil {
		return course, err
	}

	return course, nil
}

func (srv *CourseService) GetCourseList(courseMap map[string]interface{}, offset, limit int) ([]*model.CourseModel, uint64, error) {
	infos := make([]*model.CourseModel, 0)

	courses, count, err := srv.courseRepo.GetCourseList(courseMap, offset, limit)
	if err != nil {
		return nil, count, err
	}

	ids := []uint64{}
	for _, course := range courses {
		ids = append(ids, course.Id)
	}

	wg := sync.WaitGroup{}
	courseList := model.CourseList{
		Lock:  new(sync.Mutex),
		IdMap: make(map[uint64]*model.CourseModel, len(courses)),
	}

	errChan := make(chan error, 1)
	finished := make(chan bool, 1)

	// Improve query efficiency in parallel
	for _, c := range courses {
		wg.Add(1)
		go func(course *model.CourseModel) {
			defer wg.Done()

			//shortId, err := util.GenShortId()
			//if err != nil {
			//	errChan <- err
			//	return
			//}

			courseList.Lock.Lock()
			defer courseList.Lock.Unlock()

			courseList.IdMap[course.Id] = course
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
		infos = append(infos, courseList.IdMap[id])
	}

	return infos, count, nil
}

func (srv *CourseService) UpdateCourse(courseMap map[string]interface{}, id int) error {
	err := srv.courseRepo.UpdateCourse(courseMap, id)

	if err != nil {
		return err
	}

	return nil
}
