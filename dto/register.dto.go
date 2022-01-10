package dto

type RegisterDTO struct {
	OwnerName    string `json:"owner_name" form:"owner_name" validate:"required,min=3"`
	RestoName    string `json:"resto_name" form:"resto_name" validate:"required,min=3"`
	RestoSlug    string `json:"resto_slug" form:"resto_slug" validate:"required,min=3"`
	RestoAddress string `json:"resto_address" form:"resto_address"`
	Phone		string `json:"phone" form:"phone"`
	Email        string `json:"email" form:"email" validate:"required,email"`
	Password     string `json:"password" form:"password" validate:"required,min=5"`
}
