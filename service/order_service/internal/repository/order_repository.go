package repository

import (
	orderErrors "github.com/psocietyyy/go-gin-kafka/service/order_service/internal/errors"
	"github.com/psocietyyy/go-gin-kafka/service/order_service/internal/model"
	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) FindAll() ([]model.Order, error) {
	var orders []model.Order
	if err := r.db.Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *OrderRepository) Create(order *model.Order) error {
	return r.db.Create(order).Error
}

func (r *OrderRepository) FindByID(id uint) (*model.Order, error) {
	var order model.Order
	if err := r.db.First(&order, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, orderErrors.ErrOrderNotFound
		}
		return nil, err
	}
	return &order, nil
}

func (r *OrderRepository) Update(order *model.Order) error {
	return r.db.Save(order).Error
}