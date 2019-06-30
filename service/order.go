package service

import (
	"sync"
	"time"

	"github.com/1024casts/1024casts/util"

	"github.com/1024casts/1024casts/model"
	"github.com/1024casts/1024casts/repository"
)

type OrderService struct {
	orderRepo *repository.OrderRepo
}

func NewOrderService() *OrderService {
	return &OrderService{
		repository.NewOrderRepo(),
	}
}

func (srv *OrderService) CreateOrder(orderModel model.OrderModel) (uint64, error) {
	orderId, err := srv.orderRepo.CreateOrder(orderModel)
	if err != nil {
		return 0, err
	}

	return orderId, nil
}

func (srv *OrderService) GetOrderById(id int) (*model.OrderModel, error) {
	order, err := srv.orderRepo.GetOrderById(id)

	if err != nil {
		return order, err
	}

	return order, nil
}

func (srv *OrderService) GetOrderList(orderMap map[string]interface{}, offset, limit int) ([]*model.OrderInfo, uint64, error) {
	infos := make([]*model.OrderInfo, 0)

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
		IdMap: make(map[uint64]*model.OrderInfo, len(orders)),
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

			orderList.IdMap[order.Id] = srv.trans(order)
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

// 获取我的订单列表
func (srv *OrderService) GetOrderListByUserId(userId uint64, offset, limit int) ([]*model.OrderInfo, int, error) {
	infos := make([]*model.OrderInfo, 0)

	orders, count, err := srv.orderRepo.GetOrderListByUserId(userId, offset, limit)
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
		IdMap: make(map[uint64]*model.OrderInfo, len(orders)),
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

			orderList.IdMap[order.Id] = srv.trans(order)
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

func (srv *OrderService) trans(orderModel *model.OrderModel) *model.OrderInfo {
	return &model.OrderInfo{
		Id:          orderModel.Id,
		OrderId:     orderModel.OrderId,
		OrderAmount: orderModel.OrderAmount,
		PayAmount:   orderModel.PayAmount,
		PayMethod:   orderModel.PayMethod,
		PaidAt:      util.TimeToString(orderModel.PaidAt),
		CompletedAt: time.Time{},
		CanceledAt:  time.Time{},
		QrcodeId:    0,
		TradeId:     "",
		Status:      orderModel.Status,
		UserId:      orderModel.UserId,
		CreatedAt:   util.TimeToString(orderModel.CreatedAt),
		UpdatedAt:   util.TimeToString(orderModel.UpdatedAt),
		Items:       nil,
	}
}

func (srv *OrderService) UpdateComment(commentMap map[string]interface{}, id int) error {
	err := srv.orderRepo.UpdateComment(commentMap, id)

	if err != nil {
		return err
	}

	return nil
}
