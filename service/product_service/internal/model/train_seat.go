package model

import "time"

type TrainSeat struct {
	ID         uint   `gorm:"primaryKey"`
	TrainID    uint   `gorm:"not null"`
	SeatNumber string `gorm:"size:10;not null"`

	Train Train `gorm:"foreignKey:TrainID"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
