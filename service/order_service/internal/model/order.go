package model

import "time"

type OrderStatus string

const (
	OrderStatusPending OrderStatus = "PENDING"
	OrderStatusPaid    OrderStatus = "PAID"
	OrderStatusFailed  OrderStatus = "FAILED"
	OrderStatusExpired OrderStatus = "EXPIRED"
)

type Order struct {
	ID uint `gorm:"primaryKey"`
	UserID uint `gorm:"not null"`
	ScheduleSeatID uint `gorm:"not null"`
	Status OrderStatus `gorm:"type:enum('PENDING','PAID','FAILED','EXPIRED');default:'PENDING';not null"`
	
	CreatedAt time.Time
	UpdatedAt time.Time
}