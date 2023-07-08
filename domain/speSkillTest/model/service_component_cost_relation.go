package model

import "time"

type ServiceComponentCostRelation struct {
	tableName             struct{} `gorm:"service_component_type_relations"`
	Id                    int64    `gorm:"primaryKey"`
	ServiceId             int64
	ComponentCostId       int64
	ComponentCostCategory int64
	CreatedAt             time.Time `gorm:"<-create"`
	UpdatedAt             *time.Time
	DeletedAt             *time.Time
}
