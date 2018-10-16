package repository

import (
	"1024casts/backend/model"
	"1024casts/backend/pkg/constvar"
)

type OrderRepo struct {
	db *model.Database
}

func NewOrderRepo() *OrderRepo {
	return &OrderRepo{
		db: model.DB,
	}
}

func (repo *OrderRepo) GetOrderById(id int) (*model.OrderModel, error) {
	order := model.OrderModel{}
	result := repo.db.Self.Where("id = ?", id).First(&order)

	return &order, result.Error
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

func (repo *OrderRepo) UpdateComment(commentMap map[string]interface{}, id int) error {

	order, err := repo.GetOrderById(id)
	if err != nil {
		return err
	}

	return repo.db.Self.Model(order).Updates(commentMap).Error
}

func (repo *OrderRepo) DeleteComment(id int) error {
	order, err := repo.GetOrderById(id)
	if err != nil {
		return err
	}

	return repo.db.Self.Delete(&order).Error
}
