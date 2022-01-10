package dto

type RestoUpdateDTO struct {
	ID           uint64 `json:"id" form:"id"`
	OwnerName    string `json:"owner_name" form:"owner_name" validate:"omitempty,min=3"`
	RestoName    string `json:"resto_name" form:"resto_name" validate:"omitempty,min=3"`
	RestoSlug    string `json:"resto_slug" form:"resto_slug" validate:"omitempty,min=3"`
	RestoAddress string `json:"resto_address" form:"resto_address" validate:"omitempty,min=3"`
	Phone        string `json:"phone" form:"phone" validate:"omitempty,min=3"`
	Email        string `json:"email" form:"email" validate:"omitempty,email"`
	Password     string `json:"password,omitempty" form:"password,omitempty" validate:"omitempty,min=6"`
}
