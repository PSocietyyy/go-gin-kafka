package model

import "time"

type SeatStatus string

const (
	SeatStatusAvailable SeatStatus = "AVAILABLE"
	SeatStatusBooked    SeatStatus = "BOOKED"
)

type ScheduleSeat struct {
	ID          uint       `gorm:"primaryKey"`
	ScheduleID  uint       `gorm:"not null"`
	TrainSeatID uint       `gorm:"not null"`
	Status      SeatStatus `gorm:"type:enum('AVAILABLE','BOOKED');default:'AVAILABLE';not null"`

	Schedule  Schedule  `gorm:"foreignKey:ScheduleID"`
	TrainSeat TrainSeat `gorm:"foreignKey:TrainSeatID"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
