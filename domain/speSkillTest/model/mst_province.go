package model

type MstProvince struct {
	tableName struct{} `gorm:"mst_provinces"`
	Id        int64
	Name      string
}
