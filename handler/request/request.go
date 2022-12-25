package request

type UserRequest struct {
	Name     string `json:"name" validate:"required len gt 2"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required, email"`
}
