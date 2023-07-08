package model

import "time"

type QuotationAddCostInfo struct {
	tableName       struct{} `gorm:"quotation_add_cost_infos"`
	Id              int64    `gorm:"primaryKey"`
	QuotationInfoId int64
	Name            string
	Qty             *int64
	Price           *float64
	CreatedAt       time.Time `gorm:"<-create"`
	UpdatedAt       *time.Time
	DeletedAt       *time.Time
}
