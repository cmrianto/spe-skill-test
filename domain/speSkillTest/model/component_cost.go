package model

import "time"

type ComponentCost struct {
	tableName   struct{} `gorm:"component_costs"`
	Id          int64    `gorm:"primaryKey"`
	Name        string
	Description *string
	CreatedAt   time.Time `gorm:"<-create"`
	UpdatedAt   *time.Time
	DeletedAt   *time.Time
}
