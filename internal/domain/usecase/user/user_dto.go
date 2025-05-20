package user

type UserSignInRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type SignInResponse struct {
	Token string `json:"token,omitempty"`
}
