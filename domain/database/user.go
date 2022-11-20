package database

import "time"

type (
	User struct {
		ID          int       `gorm:"column:id;primary_key" json:"user_id"`
		UserName    string    `gorm:"column:user_name" json:"user_name"` // unique
		Name        string    `gorm:"column:name" json:"name"`
		Password    string    `gorm:"column:password" json:"password"`
		IDNumber    string    `gorm:"column:id_number" json:"id_number"`   // unique
		UserPhone   string    `gorm:"column:user_phone" json:"user_phone"` // unique
		UserAddress string    `gorm:"column:user_address" json:"user_address"`
		IsNew       int       `gorm:"column:is_new" json:"is_new"`
		CompanyID   int       `gorm:"column:company_id" json:"company_id"`
		UserRole    int       `gorm:"column:user_role" json:"user_role"`
		CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
		UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at"`
	}
)

func (tbl *User) TableName() string {
	return "user"
}
