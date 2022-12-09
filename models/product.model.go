package models

import "time"

type Product struct {
	ID         int64    `gorm:"primary_key:auto_increment" json:"-"`
	Title      string   `gorm:"type:text" json:"-"`
	Price      int64    `gorm:"type:bigint" json:"-"`
	Stoct      int64    `gorm:"type:bigint" json:"-"`
	CategoryID int64    `grom:"not null" json:"-"`
	Category   Category `gorm:"CategoryID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
	UserID     int64    `gorm:"not null" json:"-"`
	User       User     `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
	Created_at time.Time
	Updated_at time.Time
}
