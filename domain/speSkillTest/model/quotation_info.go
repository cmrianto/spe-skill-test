package model

import "time"

type QuotationInfo struct {
	tableName             struct{} `gorm:"quotation_infos"`
	Id                    int64    `gorm:"primaryKey"`
	QuotationId           int64
	VendorId              int64
	ServiceId             int64
	ComponentCostCategory int64
	CreatedAt             time.Time `gorm:"<-create"`
	UpdatedAt             *time.Time
	DeletedAt             *time.Time
}
