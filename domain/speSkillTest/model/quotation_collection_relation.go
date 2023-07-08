package model

import "time"

type QuotationCollectionRelation struct {
	tableName             struct{} `gorm:"quotation_collection_relations"`
	Id                    int64    `gorm:"primaryKey"`
	QuotationCollectionId int64
	QuotationId           int64
	CreatedAt             time.Time `gorm:"<-create"`
	UpdatedAt             *time.Time
	DeletedAt             *time.Time
}
