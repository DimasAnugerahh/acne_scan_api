package repository

import "acne-scan-api/internal/model/domain"

func (auth *AuthRepositoryImpl) Register(users *domain.Users) error {
	stmt, err := auth.DB.Prepare("insert into users (username,password,role,created_at,updated_at) values (?,?,?,?,?)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		users.Username,
		users.Password,
		users.Role,
		users.CreatedAt,
		users.UpdatedAt,
	)
	if err != nil {
		return err
	}

	return nil

}