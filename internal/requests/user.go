package requests

type BasicAuth struct {
	Email    string `json:"email" validate:"required,email" example:"john.doe@example.com"`
	Password string `json:"password" validate:"required,max=20,min=6" example:"11111111"`
}

type LoginRequest struct {
	BasicAuth
}

type RegisterRequest struct {
	BasicAuth
	Username string `json:"username" validate:"required" example:"John Doe"`
}

type RefreshRequest struct {
	Token string `json:"token" validate:"required" example:"refresh_token"`
}
