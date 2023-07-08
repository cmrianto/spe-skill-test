package model

import "time"

type VendorTruck struct {
	tableName struct{} `gorm:"vendor_trucks"`
	Id        int64    `gorm:"primaryKey"`
	VendorId  int64
	CreatedAt time.Time `gorm:"<-create"`
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
