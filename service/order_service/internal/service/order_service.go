package service

import (
	"github.com/psocietyyy/go-gin-kafka/service/order_service/internal/model"
	"github.com/psocietyyy/go-gin-kafka/service/order_service/internal/repository"
)

type OrderService struct {
	orderRepository *repository.OrderRepository
}

func NewOrderService(orderRepository *repository.OrderRepository) *OrderService {
	return &OrderService{orderRepository: orderRepository}
}

func (s *OrderService) FindAll() ([]model.Order, error) {
	return s.orderRepository.FindAll()
}

func (s *OrderService) Create(order *model.Order) error {
	return s.orderRepository.Create(order)
}

func (s *OrderService) FindByID(id uint) (*model.Order, error) {
	return s.orderRepository.FindByID(id)
}

func (s *OrderService) Update(order *model.Order) error {
	return s.orderRepository.Update(order)
}