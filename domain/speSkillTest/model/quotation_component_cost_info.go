package model

import "time"

type QuotationComponentCostInfo struct {
	tableName               struct{} `gorm:"quotation_component_cost_infos"`
	Id                      int64    `gorm:"primaryKey"`
	QuotationInfoId         int64
	ComponentCostId         int64
	ComponentCostLineItemId int64
	Name                    string
	Price                   *float64
	Qty                     *int64
	Note                    *string
	CreatedAt               time.Time `gorm:"<-create"`
	UpdatedAt               *time.Time
	DeletedAt               *time.Time
}
