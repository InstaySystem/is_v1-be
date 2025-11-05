package model

type Service struct {
	ID   int64  `gorm:"type:bigint;primaryKey" json:"id"`
	Name string `gorm:"type:varchar(150);not null" json:"name"`
	Slug string `gorm:"type:varchar(150);uniqueIndex:services_slug_key"`
	
}
