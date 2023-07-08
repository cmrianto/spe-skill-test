package model

import "time"

type Good struct {
	tableName         struct{} `gorm:"goods"`
	Id                int64    `gorm:"primaryKey"`
	Name              string
	UnitOfMeasurement string
	Description       *string
	CreatedAt         time.Time `gorm:"<-create"`
	UpdatedAt         *time.Time
	DeletedAt         *time.Time
}
