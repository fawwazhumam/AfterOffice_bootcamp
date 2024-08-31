package request

type ShopCreateRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Address  string `json:"address"`
	Contact  string `json:"contact"`
}

type ShopUpdateRequest struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Contact string `json:"contact"`
}

type ShopEmailRequest struct {
	Email string `json:"email" validate:"required"`
}