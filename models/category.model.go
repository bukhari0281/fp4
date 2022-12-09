package models

import "time"

type Category struct {
	ID        int64  `gorm:"primary_key:auto_increment" json:"-"`
	Type      string `gorm:"type:text" json:"-"`
	Spa       int64  `gorm:"type:bigint" json:"-"`
	ProductID int64  `gorm:"not null" json:"-"`
	// Products   []Product `gorm:"many2many:ProductID" json:"-"`
	UserID     int64 `gorm:"not null" json:"-"`
	User       User  `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
	Created_at time.Time
	Updated_at time.Time
}
