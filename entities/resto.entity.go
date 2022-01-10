package entities

type Resto struct {
	ID           uint64 `gorm:"primary_key:auto_increment" json:"id"`
	OwnerName    string `gorm:"type:varchar(255)" json:"owner_name"`
	RestoName    string `gorm:"type:varchar(255)" json:"resto_name"`
	RestoSlug    string `gorm:"type:varchar(255);uniqueIndex" json:"resto_slug"`
	RestoAddress string `gorm:"type:varchar(255)" json:"resto_address"`
	Phone        string `gorm:"type:varchar(255)" json:"phone"`
	Email        string `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Password     string `gorm:"->;<-;not null" json:"-"`
	Token        string `gorm:"-" json:"token,omitempty"`
}
