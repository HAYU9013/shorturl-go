package models

type URLMapping struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	ShortCode string `json:"short_code" gorm:"unique"`
	LongURL   string `json:"long_url"`
}
