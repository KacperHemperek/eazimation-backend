package services

import (
	"database/sql"
	"eazimation-backend/internal/database"
)

func scanUser(row sql.Row) (*database.UserModel, error) {
	user := &database.UserModel{}
	err := row.Scan(&user.ID, &user.Email, &user.Avatar)
	if err != nil {
		return nil, err
	}
	return user, nil
}
