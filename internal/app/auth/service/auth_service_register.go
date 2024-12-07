package service

import (
	"acne-scan-api/internal/model/web"
	conversion "acne-scan-api/internal/pkg/conversion/request"
	"fmt"
	"time"
)

func (auth *AuthServiceImpl) Register(request *web.Register) error {
	var err error

	err = auth.Validator.Struct(request)
	if err != nil {
		return err
	}

	users := conversion.RegisterToUserModel(request)

	wib, err := time.LoadLocation("Asia/Jakarta") // WIB (UTC+7)
	if err != nil {
		return fmt.Errorf("error loading WIB location: %s", err.Error())
	}

	createdAt := time.Now().In(wib)
	users.CreatedAt = createdAt
	users.UpdatedAt = createdAt

	err = auth.AuthRepository.Register(users)
	if err != nil {
		return fmt.Errorf("error register %s", err.Error())
	}

	return nil
}
