package model

import "time"

type User struct {
	ID        int64     `gorm:"type:bigint;primaryKey" json:"id"`
	Username  string    `gorm:"type:varchar(50);uniqueIndex:users_username_key;not null" json:"username"`
	Email     string    `gorm:"type:varchar(150);uniqueIndex:users_email_key;not null" json:"email"`
	Role      string    `gorm:"type:varchar(20);check:role IN ('receptionist', 'housekeeper', 'technician', 'admin')" json:"role"`
	FirstName string    `gorm:"type:varchar(150);not null" json:"first_name"`
	LastName  string    `gorm:"type:varchar(150);not null" json:"last_name"`
	Phone     string    `gorm:"type:char(10);uniqueIndex:users_phone_key;not null" json:"phone"`
	Password  string    `gorm:"type:varchar(255);not null" json:"password"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
