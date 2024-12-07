package conversion

import (
	"acne-scan-api/internal/model/domain"
	"acne-scan-api/internal/model/web"
)

func RegisterToUserModel(req *web.Register)*domain.Users{
	return &domain.Users{
		Username: req.Username,
		Password: req.Password,
		Role: req.Role,
	}
}