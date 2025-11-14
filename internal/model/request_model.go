package model

import "time"

type RequestType struct {
	ID   int64  `gorm:"type:bigint;primaryKey" json:"id"`
	Name string `gorm:"type:varchar(150);not null" json:"name"`
	Slug string `gorm:"type:varchar(150);uniqueIndex:services_slug_key"`
}

type Request struct {
	ID          int64     `gorm:"type:bigint;primaryKey" json:"id"`
	Content     string    `gorm:"type:text;not null" json:"content"`
	Status      string    `gorm:"type:varchar(20);check:status IN ('pending', 'accepted', 'canceled', 'done')" json:"status"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	UpdatedByID int64     `gorm:"type:bigint;not null" json:"updated_by_id"`

	UpdatedBy *User `gorm:"foreignKey:UpdatedByID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT" json:"updated_by"`
}
