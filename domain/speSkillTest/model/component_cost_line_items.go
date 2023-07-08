package model

import "time"

type ComponentCostLineItem struct {
	tableName       struct{} `gorm:"component_cost_line_items"`
	Id              int64    `gorm:"primaryKey"`
	ComponentCostId int64
	Name            string
	Price           float64
	Qty             int64
	Notes           *string
	CreatedAt       time.Time `gorm:"<-create"`
	UpdatedAt       *time.Time
	DeletedAt       *time.Time
}
