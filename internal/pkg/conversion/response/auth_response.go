package response

import "acne-scan-api/internal/model/web"

func AuthResponse(authRequest web.UserLogin, token string,role string) *web.AuthResponse {
	return &web.AuthResponse{
		Username: authRequest.Username,
		Token:    token,
		Role:     role,
	}
}
