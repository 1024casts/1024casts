package service

import (
	"strconv"
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

const (
	UserMemberStatusNormal = 1 // 正常
	UserMemberStatusDelete = 2 // 删除

	// 会员类型
	UserMemberTypeMonth     = 1 // 一个月
	UserMemberTypeQuarter   = 2 // 3个月
	UserMemberTypeHalfYear  = 3 // 半年
	UserMemberTypeYear      = 4 // 一年
	UserMemberTypeTwoYear   = 5 // 2年
	UserMemberTypeThreeYear = 6 // 3年

	UserMemberTypeTest = 15 // 测试商品 1个月
)

type OrderService struct {
	orderRepo *repository.OrderRepo
	userRepo  *repository.UserRepo
}

func NewOrderService() *OrderService {
	return &OrderService{
		orderRepo: repository.NewOrderRepo(),
		userRepo:  repository.NewUserRepo(),
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
	resp, err := payNative.Create(int64(payAmount), plan.Name, strconv.Itoa(int(orderId)), "", "")
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

func (srv *OrderService) GetOrderItemById(orderId int) (*model.OrderItemModel, error) {
	orderItem, err := srv.orderRepo.GetOrderItemById(orderId)

	if err != nil {
		return orderItem, err
	}

	return orderItem, nil
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

// 确认订单已支付
func (srv *OrderService) ConfirmOrderPaid(order *model.OrderModel, orderItem *model.OrderItemModel, payTime string) (err error) {

	db := model.DB.Self
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			log.Warnf("[order] recover begin tx err: %+v", err)
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		log.Warnf("[order] begin tx err: %+v", err)
		return err
	}

	orderMap := map[string]interface{}{
		"status": constvar.OrderStatusPaid,
		"PaidAt": payTime,
	}
	err = srv.orderRepo.UpdateStatus(tx, order.Id, orderMap)
	if err != nil {
		tx.Rollback()
		log.Warnf("[order] update order status err: %+v", err)
		return err
	}

	// 给用户写入会员信息 user_members
	userMember, err := srv.userRepo.GetUserMember(order.UserId, UserMemberStatusNormal)
	if err != nil {
		tx.Rollback()
		log.Warnf("[order] get user member err: %+v", err)
		return err
	}

	startTime, err := util.StringToTime(payTime)
	if err != nil {
		tx.Rollback()
		log.Warnf("[order] string to time err: %+v", err)
		return err
	}
	// 如果不存在，则为首次购买会员
	if userMember.Id == 0 {
		userMemberModel := model.UserMemberModel{
			EndTime:   srv.getUserMemberEndTime(startTime, orderItem.ItemID),
			StartTime: startTime,
			Status:    UserMemberStatusNormal,
			Type:      orderItem.ItemID, // 目前只有获取结束时间有用
			UserID:    order.UserId,
		}
		if err := tx.Create(&userMemberModel).Error; err != nil {
			tx.Rollback()
			log.Warnf("[order] create user member err: %v", err)
			return err
		}
	} else {
		// 续期
		// a. 如果会员还未到期就续费, 则在结束时间上再加对应的时间段即可
		if time.Now().Unix() <= userMember.EndTime.Unix() {
			// 更新到期时间为最新
			endTime := srv.getUserMemberEndTime(startTime, orderItem.ItemID)
			err = srv.userRepo.UpdateUserMemberEndTime(tx, order.UserId, endTime)
			if err != nil {
				tx.Rollback()
				log.Warnf("[order] update user member end time err: %v", err)
				return err
			}
		} else {
			// b. 如果购买的会员已经过期再次购买, 则直接更新现有记录为最新的会员时间
			endTime := srv.getUserMemberEndTime(startTime, orderItem.ItemID)
			err = srv.userRepo.UpdateUserMember(tx, order.UserId, startTime, endTime)
			if err != nil {
				tx.Rollback()
				log.Warnf("[order] update user member err: %v", err)
				return err
			}
		}
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		log.Warnf("[order] tx commit err: %+v", err)
		return err
	}

	return nil
}

// 获取会员的结束时间
func (srv *OrderService) getUserMemberEndTime(startTime time.Time, level uint64) time.Time {
	endTime := time.Time{}
	// 一个月按31天算
	monthDuration := 24 * 3600 * 31 * time.Second
	switch level {
	// 一个月
	case UserMemberTypeMonth:
		endTime = startTime.Add(monthDuration)
		break
	// 3个月
	case UserMemberTypeQuarter:
		endTime = startTime.Add(monthDuration * 3)
		break
	// 半年
	case UserMemberTypeHalfYear:
		endTime = startTime.Add(monthDuration * 6)
		break
	// 一年
	case UserMemberTypeYear:
		endTime = startTime.Add(monthDuration * 12)
		break
	// 2年
	case UserMemberTypeTwoYear:
		endTime = startTime.Add(monthDuration * 24 * 2)
		break
	// 3年
	case UserMemberTypeThreeYear:
		endTime = startTime.Add(monthDuration * 24 * 3)
		break
	// 测试商品
	case UserMemberTypeTest:
		endTime = startTime.Add(monthDuration)
		break
	}

	return endTime
}
