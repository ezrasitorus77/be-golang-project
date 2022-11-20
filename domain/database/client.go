package database

import "time"

type (
	Client struct {
		ID            int       `gorm:"column:id;primary_key" json:"client_id"`
		ClientName    string    `gorm:"column:client_name" json:"client_name"` // unique
		ClientParent  string    `gorm:"column:client_parent" json:"client_parent"`
		ClientField   int       `gorm:"column:client_field" json:"client_field"`
		ClientAddress string    `gorm:"column:client_address" json:"client_address"`
		ClientPhone   string    `gorm:"column:client_phone" json:"client_phone"`     // unique
		ClientWebsite string    `gorm:"column:client_website" json:"vendor_website"` // unique
		Email         string    `gorm:"column:client_email" json:"client_email"`     // unique
		SocialMedia   string    `gorm:"column:social_media" json:"social_media"`
		Province      string    `gorm:"column:province" json:"province"`
		City          string    `gorm:"column:city" json:"city"`
		District      string    `gorm:"column:district" json:"district"`
		Avatar        string    `gorm:"column:avatar" json:"avatar"`
		IsNew         int       `gorm:"column:is_new" json:"is_new"`
		CreatedAt     time.Time `gorm:"column:created_at" json:"created_at"`
		UpdatedAt     time.Time `gorm:"column:updated_at" json:"updated_at"`
	}
)

func (tbl *Client) TableName() string {
	return "client"
}
