package model

import "time"

type Schedule struct {
	ID            uint      `gorm:"primaryKey"`
	TrainID       uint      `gorm:"not null"`
	DepartureTime time.Time `gorm:"not null"`
	ArrivalTime   time.Time `gorm:"not null"`

	Train Train `gorm:"foreignKey:TrainID"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
