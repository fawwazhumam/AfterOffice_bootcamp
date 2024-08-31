package request

type ShopRequest struct {

	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}