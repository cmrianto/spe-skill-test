package model

import "time"

type Service struct {
	tableName   struct{} `gorm:"services"`
	Id          int64    `gorm:"primaryKey"`
	Name        string
	Description *string
	CreatedAt   time.Time `gorm:"<-create"`
	UpdatedAt   *time.Time
	DeletedAt   *time.Time
}
