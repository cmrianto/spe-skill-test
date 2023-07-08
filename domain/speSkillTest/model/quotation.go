package model

import "time"

type Quotation struct {
	tableName                  struct{} `gorm:"quotations"`
	Id                         int64    `gorm:"primaryKey"`
	QuotationNumber            string
	Date                       *time.Time
	ServiceId                  int64
	Note                       *string
	ScheduleId                 *int64
	ScheduleEtd                *time.Time
	ScheduleEta                *time.Time
	ScheduleClosingTime        *time.Time
	ClientId                   *int64
	ClientWarehouseId          *int64
	ClientWarehouseAddress     *string
	ConsigneeId                *int64
	ConsigneeWarehouseId       *int64
	ConsigneeWarehouseAddress  *string
	PortOfLoadingId            *int64
	PortOfDestinationId        *int64
	GoodId                     *int64
	GoodUnitOfMeasurement      *string
	GoodQty                    *int64
	ContainerId                *int64
	ContainerUnitOfMeasurement *string
	ContainerNote              *string
	CommissionAType            *int64
	CommissionATypeValue       *float64
	CommissionAValue           *float64
	CommissionANote            *string
	CommissionBType            *int64
	CommissionBTypeValue       *float64
	CommissionBValue           *float64
	CommissionBNote            *string
	MarginType                 *int64
	MarginTypeValue            *float64
	MarginValue                *float64
	MarginRoundValue           *float64
	MarginNote                 *string
	FinalPriceValue            *float64
	FinalPriceNote             *string
	CreatedAt                  time.Time `gorm:"<-create"`
	UpdatedAt                  *time.Time
	DeletedAt                  *time.Time
}
