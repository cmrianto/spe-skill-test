package model

import "time"

type QuotationCollection struct {
	tableName  struct{} `gorm:"quotation_collections"`
	Id         int64    `gorm:"primaryKey"`
	ClientName string
	CreatedAt  time.Time `gorm:"<-create"`
	UpdatedAt  *time.Time
	DeletedAt  *time.Time
}
