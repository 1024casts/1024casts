package repository

import (
	"github.com/1024casts/1024casts/model"
	"github.com/1024casts/1024casts/pkg/constvar"
)

type CourseRepo struct {
	db *model.Database
}

func NewCourseRepo() *CourseRepo {
	return &CourseRepo{
		db: model.DB,
	}
}

func (repo *CourseRepo) CreateCourse(course model.CourseModel) (id uint64, err error) {
	err = repo.db.Self.Create(&course).Error
	if err != nil {
		return 0, err
	}

	return course.Id, nil
}

func (repo *CourseRepo) GetCourseById(id int) (*model.CourseModel, error) {
	course := model.CourseModel{}
	result := repo.db.Self.Where("id = ?", id).First(&course)

	return &course, result.Error
}

func (repo *CourseRepo) GetCourseBySlug(slug string) (*model.CourseModel, error) {
	course := model.CourseModel{}
	result := repo.db.Self.Where("slug = ?", slug).First(&course)

	return &course, result.Error
}

func (repo *CourseRepo) GetCourseList(courseMap map[string]interface{}, offset, limit int) ([]*model.CourseModel, int, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	courses := make([]*model.CourseModel, 0)
	var count int

	if err := repo.db.Self.Model(&model.CourseModel{}).Where(courseMap).Count(&count).Error; err != nil {
		return courses, count, err
	}

	if err := repo.db.Self.Where(courseMap).Offset(offset).Limit(limit).Order("id desc").Find(&courses).Error; err != nil {
		return courses, count, err
	}

	return courses, count, nil
}

func (repo *CourseRepo) GetSectionList(courseId uint64) ([]*model.SectionModel, error) {
	sections := make([]*model.SectionModel, 0)

	if err := repo.db.Self.Where("course_id = ?", courseId).Order("weight asc").Find(&sections).Error; err != nil {
		return sections, err
	}

	return sections, nil
}

func (repo *CourseRepo) UpdateCourse(userMap map[string]interface{}, id int) error {

	course, err := repo.GetCourseById(id)
	if err != nil {
		return err
	}

	return repo.db.Self.Model(course).Updates(userMap).Error
}

func (repo *CourseRepo) DeleteCourse(id int) error {
	course, err := repo.GetCourseById(id)
	if err != nil {
		return err
	}

	return repo.db.Self.Delete(&course).Error
}

func (repo *CourseRepo) Store(course *model.CourseModel) (id uint64, err error) {
	//users := model.CourseModel{}

	return 0, nil
}
