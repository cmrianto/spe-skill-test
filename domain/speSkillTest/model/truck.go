package model

import "time"

type Truck struct {
	tableName     struct{} `gorm:"trucks"`
	Id            int64    `gorm:"primaryKey"`
	VendorTruckId int64
	ProvinceId    int64
	Name          string
	Price         float64
	CreatedAt     time.Time `gorm:"<-create"`
	UpdatedAt     *time.Time
	DeletedAt     *time.Time
}
