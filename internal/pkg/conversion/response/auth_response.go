package response

import "acne-scan-api/internal/model/web"

func AuthResponse(authRequest web.UserLogin, token string, role string, userId int) *web.AuthResponse {
	return &web.AuthResponse{
		Username: authRequest.Username,
		UserId:   userId,
		Token:    token,
		Role:     role,
	}
}
