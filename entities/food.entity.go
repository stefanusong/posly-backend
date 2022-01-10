package entities

type Food struct {
	ID         uint64  `gorm:"primary_key:auto_increment" json:"id"`
	Name       string  `gorm:"type:varchar(255)" json:"name"`
	Desc       string  `gorm:"type:varchar(255)" json:"desc"`
	CategoryId uint64  `gorm:"type:varchar(255)" json:"category_id"`
	Price      float64 `gorm:"type:varchar(255)" json:"price"`
	Stock      int     `gorm:"type:varchar(255)" json:"stock"`
}
