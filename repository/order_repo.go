package repository

import (
	"github.com/1024casts/1024casts/model"
	"github.com/1024casts/1024casts/pkg/constvar"
	"github.com/jinzhu/gorm"
)

type OrderRepo struct {
	db *model.Database
}

func NewOrderRepo() *OrderRepo {
	return &OrderRepo{
		db: model.DB,
	}
}

func (repo *OrderRepo) CreateOrder(order model.OrderModel) (id uint64, err error) {
	err = repo.db.Self.Create(&order).Error
	if err != nil {
		return 0, err
	}

	return order.Id, nil
}

func (repo *OrderRepo) GetOrderById(id int) (*model.OrderModel, error) {
	order := model.OrderModel{}
	result := repo.db.Self.Where("id = ?", id).First(&order)

	return &order, result.Error
}

func (repo *OrderRepo) GetOrderItemById(orderId int) (*model.OrderItemModel, error) {
	orderItem := model.OrderItemModel{}
	result := repo.db.Self.Where("order_id = ?", orderId).First(&orderItem)

	return &orderItem, result.Error
}

func (repo *OrderRepo) GetOrderList(orderMap map[string]interface{}, offset, limit int) ([]*model.OrderModel, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	orders := make([]*model.OrderModel, 0)
	var count uint64

	if err := repo.db.Self.Model(&model.OrderModel{}).Where(orderMap).Count(&count).Error; err != nil {
		return orders, count, err
	}

	if err := repo.db.Self.Where(orderMap).Offset(offset).Limit(limit).Order("id desc").Find(&orders).Error; err != nil {
		return orders, count, err
	}

	return orders, count, nil
}

func (repo *OrderRepo) GetOrderListByUserId(userId uint64, offset, limit int) ([]*model.OrderModel, int, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	orders := make([]*model.OrderModel, 0)
	var count int

	if err := repo.db.Self.Model(&model.OrderModel{}).Where("user_id=?", userId).Count(&count).Error; err != nil {
		return orders, count, err
	}

	if err := repo.db.Self.Where("user_id=?", userId).Offset(offset).Limit(limit).Order("id desc").
		Preload("Items").Find(&orders).Error; err != nil {
		return orders, count, err
	}

	return orders, count, nil
}

// update order status
func (repo *OrderRepo) UpdateStatus(db *gorm.DB, id uint64, orderMap map[string]interface{}) error {

	order := model.OrderModel{}

	return db.Model(order).Where("id=?", id).Updates(orderMap).Error
}
