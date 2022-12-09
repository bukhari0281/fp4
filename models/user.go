package models

type User struct {
	ID        int64  `gorm:"primary_key:auto_increment" json:"-"`
	Full_name string `gorm:"type:varchar(100)" json:"-"`
	Email     string `gorm:"type:varchar(100);unique;" json:"-"`
	Password  string `gorm:"type:varchar(100)" json:"-"`
	Role      string `gorm:"type:varchar(100)" json:"-"`
	Balance   int64  `gorm:"type:bigint" json:"-"`
}
