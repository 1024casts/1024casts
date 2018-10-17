package repository

import (
	"1024casts/backend/model"
	"1024casts/backend/pkg/constvar"
)

type PlanRepo struct {
	db *model.Database
}

func NewPlanRepo() *PlanRepo {
	return &PlanRepo{
		db: model.DB,
	}
}

func (repo *PlanRepo) CreatePlan(plan model.PlanModel) (id uint64, err error) {
	err = repo.db.Self.Create(&plan).Error
	if err != nil {
		return 0, err
	}

	return plan.ID, nil
}

func (repo *PlanRepo) GetPlanById(id int) (*model.PlanModel, error) {
	plan := model.PlanModel{}
	result := repo.db.Self.Where("id = ?", id).First(&plan)

	return &plan, result.Error
}

func (repo *PlanRepo) GetPlanByAlias(alias string) (*model.PlanModel, error) {
	plan := model.PlanModel{}
	result := repo.db.Self.Where("alias = ?", alias).First(&plan)

	return &plan, result.Error
}

func (repo *PlanRepo) GetPlanList(planMap map[string]interface{}, offset, limit int) ([]*model.PlanModel, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	plans := make([]*model.PlanModel, 0)
	var count uint64

	if err := repo.db.Self.Model(&model.PlanModel{}).Where(planMap).Count(&count).Error; err != nil {
		return plans, count, err
	}

	if err := repo.db.Self.Where(planMap).Offset(offset).Limit(limit).Order("id desc").Find(&plans).Error; err != nil {
		return plans, count, err
	}

	return plans, count, nil
}

func (repo *PlanRepo) UpdatePlan(userMap map[string]interface{}, id int) error {

	plan, err := repo.GetPlanById(id)
	if err != nil {
		return err
	}

	return repo.db.Self.Model(plan).Updates(userMap).Error
}

func (repo *PlanRepo) DeletePlan(id int) error {
	plan, err := repo.GetPlanById(id)
	if err != nil {
		return err
	}

	return repo.db.Self.Delete(&plan).Error
}

func (repo *PlanRepo) Store(course *model.CourseModel) (id uint64, err error) {
	//users := model.CourseModel{}

	return 0, nil
}
