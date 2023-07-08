package model

import "time"

type QuotationTruckInfo struct {
	tableName       struct{} `gorm:"quotation_truck_infos"`
	Id              int64    `gorm:"primaryKey"`
	QuotationInfoId int64
	VendorTruckId   int64
	TruckId         int64
	Price           *float64
	Qty             *int64
	Note            *string
	CreatedAt       time.Time `gorm:"<-create"`
	UpdatedAt       *time.Time
	DeletedAt       *time.Time
}
