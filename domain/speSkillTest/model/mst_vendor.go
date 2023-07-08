package model

import "time"

type MstVendor struct {
	tableName     struct{} `gorm:"mst_vendors"`
	Id            int64    `gorm:"primaryKey"`
	ProvinceId    *int64
	CityId        *int64
	DistrictId    *int64
	CompanyName   string
	Email         string
	Npwp          string
	StreetAddress string
	CreatedAt     time.Time `gorm:"<-create"`
	UpdatedAt     *time.Time
	DeletedAt     *time.Time
}
