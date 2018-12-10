package service

import (
	"sync"

	"github.com/1024casts/1024casts/model"
	"github.com/1024casts/1024casts/repository"
	"github.com/1024casts/1024casts/util"

	"github.com/lexkong/log"
)

type CourseService struct {
	repo     *repository.CourseRepo
	videoSrv *VideoService
}

func NewCourseService() *CourseService {
	return &CourseService{
		repo:     repository.NewCourseRepo(),
		videoSrv: NewVideoService(),
	}
}

func (srv *CourseService) CreateCourse(user model.CourseModel) (id uint64, err error) {
	id, err = srv.repo.CreateCourse(user)

	if err != nil {
		return id, err
	}

	return id, nil
}

func (srv *CourseService) GetCourseById(id int) (*model.CourseModel, error) {
	course, err := srv.repo.GetCourseById(id)

	if err != nil {
		return course, err
	}

	return course, nil
}

func (srv *CourseService) GetCourseList(courseMap map[string]interface{}, offset, limit int) ([]*model.CourseModel, uint64, error) {
	infos := make([]*model.CourseModel, 0)

	courses, count, err := srv.repo.GetCourseList(courseMap, offset, limit)
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

			course.CoverImage = util.GetQiniuPrivateAccessUrl(course.CoverImage)
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

func (srv *CourseService) GetCourseSectionList(courseId uint64, offset, limit int) ([]*model.SectionModel, uint64, error) {
	infos := make([]*model.SectionModel, 0)

	sections, count, err := srv.repo.GetSectionList(courseId, offset, limit)
	if err != nil {
		return nil, count, err
	}

	ids := []uint64{}
	for _, section := range sections {
		ids = append(ids, section.Id)
	}

	wg := sync.WaitGroup{}
	sectionList := model.SectionList{
		Lock:  new(sync.Mutex),
		IdMap: make(map[uint64]*model.SectionModel, len(sections)),
	}

	videos, err := srv.videoSrv.GetVideoList(courseId)
	if err != nil {
		log.Warnf("[course] get video list fail from video repo, course_id: %d", courseId)
	}

	errChan := make(chan error, 1)
	finished := make(chan bool, 1)

	// Improve query efficiency in parallel
	for _, c := range sections {
		wg.Add(1)
		go func(section *model.SectionModel) {
			defer wg.Done()

			sectionList.Lock.Lock()
			defer sectionList.Lock.Unlock()

			for _, video := range videos {
				if section.Id == video.SectionID {
					section.VideoItems = append(section.VideoItems, video)
				}
			}

			sectionList.IdMap[section.Id] = section
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
		infos = append(infos, sectionList.IdMap[id])
	}

	return infos, count, nil
}

func (srv *CourseService) UpdateCourse(courseMap map[string]interface{}, id int) error {
	err := srv.repo.UpdateCourse(courseMap, id)

	if err != nil {
		return err
	}

	return nil
}
