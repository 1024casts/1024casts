package service

import (
	"sync"

	"1024casts/backend/model"
	"1024casts/backend/repository"
	"strconv"
)

type OrderService struct {
	orderRepo *repository.OrderRepo
}

func NewOrderService() *OrderService {
	return &OrderService{
		repository.NewOrderRepo(),
	}
}

func (srv *OrderService) GetOrderById(id int) (*model.OrderModel, error) {
	order, err := srv.orderRepo.GetOrderById(id)

	if err != nil {
		return order, err
	}

	return order, nil
}

func (srv *OrderService) GetOrderList(orderMap map[string]interface{}, offset, limit int) ([]*model.OrderModel, uint64, error) {
	infos := make([]*model.OrderModel, 0)

	orders, count, err := srv.orderRepo.GetOrderList(orderMap, offset, limit)
	if err != nil {
		return nil, count, err
	}

	ids := []uint64{}
	for _, order := range orders {
		ids = append(ids, order.Id)
	}

	wg := sync.WaitGroup{}
	orderList := model.OrderList{
		Lock:  new(sync.Mutex),
		IdMap: make(map[uint64]*model.OrderModel, len(orders)),
	}

	errChan := make(chan error, 1)
	finished := make(chan bool, 1)

	// Improve query efficiency in parallel
	for _, o := range orders {
		wg.Add(1)
		go func(order *model.OrderModel) {
			defer wg.Done()

			//shortId, err := util.GenShortId()
			//if err != nil {
			//	errChan <- err
			//	return
			//}

			orderList.Lock.Lock()
			defer orderList.Lock.Unlock()

			order.OrderId = strconv.FormatInt(int64(order.Id), 10)
			orderList.IdMap[order.Id] = order
		}(o)
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
		infos = append(infos, orderList.IdMap[id])
	}

	return infos, count, nil
}

func (srv *OrderService) UpdateComment(commentMap map[string]interface{}, id int) error {
	err := srv.orderRepo.UpdateComment(commentMap, id)

	if err != nil {
		return err
	}

	return nil
}
