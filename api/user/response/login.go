package response

type LoginResponse struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

func NewLoginResponse(email string, token string) *LoginResponse {
	var loginResponse LoginResponse
	loginResponse.Email = email
	loginResponse.Token = token
	return &loginResponse
}
