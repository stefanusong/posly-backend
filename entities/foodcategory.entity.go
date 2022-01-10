package entities

type FoodCategory struct {
	ID      uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Name    string `gorm:"type:varchar(255)" json:"name"`
	RestoId uint64 `gorm:"type:varchar(255)" json:"resto_id"`
}
