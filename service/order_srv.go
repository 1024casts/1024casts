package service

import (
	"sync"
	"time"

	"github.com/1024casts/1024casts/pkg/constvar"

	"github.com/qingwg/payjs"
	"github.com/qingwg/payjs/native"
	"github.com/spf13/viper"

	"github.com/1024casts/1024casts/util"
	"github.com/lexkong/log"

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

// 创建订单
func (srv *OrderService) CreateOrder(userId uint64, plan *model.PlanModel) (id uint64, qrCodeUrl string, err error) {

	db := model.DB.Self
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		return 0, "", err
	}

	orderId := util.GenOrderNo()
	orderInput := model.OrderModel{
		Id:          orderId,
		PayMethod:   constvar.PayMethodWeiXin,
		OrderAmount: plan.Price,
		PayAmount:   plan.Price,
		UserId:      userId,
		Status:      constvar.OrderStatusPending,
	}
	// step1: create main order
	if err := tx.Create(&orderInput).Error; err != nil {
		tx.Rollback()
		log.Warnf("[order] create user err: %v", err)
		return 0, "", err
	}

	// step2: create pay order
	payResp, err := srv.CreatePayOrder(orderId, plan)
	if err != nil {
		tx.Rollback()
		log.Warnf("[order] create user pay order err: %v", err)
		return 0, "", err
	}

	// step3: write to order items
	orderItemModel := model.OrderItemModel{
		OrderID:  orderId,
		UserID:   userId,
		ItemID:   plan.ID,
		Name:     plan.Name,
		Price:    plan.Price,
		Quantity: 1,
		Amount:   plan.Price * 1,
	}
	if err := tx.Create(&orderItemModel).Error; err != nil {
		tx.Rollback()
		log.Warnf("[order] create order item err: %v", err)
		return 0, "", err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		log.Warnf("[order] commit order fail, err: %+v", err)
		return 0, "", err
	}

	return orderId, payResp.Qrcode, nil
}

// 调用第三方服务创建支付订单
func (srv *OrderService) CreatePayOrder(orderId uint64, plan *model.PlanModel) (native.CreateResponse, error) {
	payJsConfig := &payjs.Config{
		Key:       viper.GetString("pay_js.secret_key"),
		MchID:     viper.GetString("pay_js.mch_id"),
		NotifyUrl: "",
	}
	pay := payjs.New(payJsConfig)
	payNative := pay.GetNative()
	payAmount := plan.Price * 100
	resp, err := payNative.Create(int64(payAmount), plan.Name, string(orderId), "", "")
	if err != nil {
		log.Warnf("[plans] create pay order err: %v", err)
		return native.CreateResponse{}, err
	}
	log.Infof("[plans] pay create resp: %+v", resp)
	return resp, nil
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
		PayMethod:   util.GetPayMethodText(orderModel.PayMethod),
		PaidAt:      util.TimeToString(orderModel.PaidAt),
		CompletedAt: time.Time{},
		CanceledAt:  time.Time{},
		QrcodeId:    0,
		TradeId:     "",
		Status:      orderModel.Status,
		UserId:      orderModel.UserId,
		CreatedAt:   util.TimeToString(orderModel.CreatedAt),
		UpdatedAt:   util.TimeToString(orderModel.UpdatedAt),
		OrderItems:  orderModel.Items,
	}
}
