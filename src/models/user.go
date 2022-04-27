package models

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type UserShort struct {
	FirstName string `json:"first_name" binding:"required,lte=255"`
	Username  string `json:"username" binding:"required,lte=255"`
	Email     string `json:"email" binding:"required,email,lte=320"`
}

type UserFull struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	FirstName string    `json:"first_name"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
}

func CreateNewUser(newUser UserShort) (UserFull, error) {
	id := uuid.New()
	var createdUser UserFull
	err := db.QueryRow(context.Background(),
		"INSERT INTO blog_user(id, first_name, username, email) VALUES($1, $2, $3, $4) RETURNING id, created_at, first_name, username, email",
		id, newUser.FirstName, newUser.Username, newUser.Email).Scan(
		&createdUser.ID,
		&createdUser.CreatedAt,
		&createdUser.FirstName,
		&createdUser.Username,
		&createdUser.Email)
	if err != nil {
		return createdUser, err
	}
	return createdUser, nil
}

func GetSlice(orderBy string) ([]UserFull, error) {
	usersSlice := make([]UserFull, 0)
	rows, err := db.Query(context.Background(), "SELECT * FROM blog_user ORDER BY ($1)", orderBy)
	if err != nil {
		return usersSlice, err
	}
	next := rows.Next()
	if !next {
		return usersSlice, nil
	}
	for next {
		var user UserFull
		rows.Scan(
			&user.ID,
			&user.CreatedAt,
			&user.FirstName,
			&user.Username,
			&user.Email)
		usersSlice = append(usersSlice, user)
		next = rows.Next()
	}
	return usersSlice, nil
}
