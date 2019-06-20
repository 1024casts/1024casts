package service

import (
	"sync"

	"github.com/1024casts/1024casts/model"
	"github.com/1024casts/1024casts/repository"
	"github.com/1024casts/1024casts/util"
)

type PlanService struct {
	repo *repository.PlanRepo
}

func NewPlanService() *PlanService {
	return &PlanService{
		repository.NewPlanRepo(),
	}
}

func (srv *PlanService) CreateCourse(plan model.PlanModel) (id uint64, err error) {
	id, err = srv.repo.CreatePlan(plan)

	if err != nil {
		return id, err
	}

	return id, nil
}

func (srv *PlanService) GetPlanById(id int) (*model.PlanModel, error) {
	plan, err := srv.repo.GetPlanById(id)

	if err != nil {
		return plan, err
	}

	return plan, nil
}

func (srv *PlanService) GetPlanByAlias(alias string) (*model.PlanModel, error) {
	plan, err := srv.repo.GetPlanByAlias(alias)

	if err != nil {
		return plan, err
	}

	return plan, nil
}

func (srv *PlanService) GetPlanList(courseMap map[string]interface{}, offset, limit int) ([]*model.PlanModel, uint64, error) {
	infos := make([]*model.PlanModel, 0)

	plans, count, err := srv.repo.GetPlanList(courseMap, offset, limit)
	if err != nil {
		return nil, count, err
	}

	ids := []uint64{}
	for _, plan := range plans {
		ids = append(ids, plan.ID)
	}

	wg := sync.WaitGroup{}
	planList := model.PlanList{
		Lock:  new(sync.Mutex),
		IdMap: make(map[uint64]*model.PlanModel, len(plans)),
	}

	errChan := make(chan error, 1)
	finished := make(chan bool, 1)

	// Improve query efficiency in parallel
	for _, c := range plans {
		wg.Add(1)
		go func(plan *model.PlanModel) {
			defer wg.Done()

			planList.Lock.Lock()
			defer planList.Lock.Unlock()

			plan.FormatCreatedAt = util.TimeToString(plan.CreatedAt)
			plan.FormatUpdatedAt = util.TimeToString(plan.UpdatedAt)
			planList.IdMap[plan.ID] = plan
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
		infos = append(infos, planList.IdMap[id])
	}

	return infos, count, nil
}

func (srv *PlanService) UpdatePlan(courseMap map[string]interface{}, id int) error {
	err := srv.repo.UpdatePlan(courseMap, id)

	if err != nil {
		return err
	}

	return nil
}
