package model

import "time"

type Train struct {
	ID		uint `gorm:"primaryKey"`
	Name	string	`gorm:"size:100;not null"`
	Code	string	`gorm:"uniqueIndex;size:20;not null"`

	
	
	CreatedAt	time.Time
	UpdatedAt	time.Time
}