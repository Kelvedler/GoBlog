package models

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
)

type UserShort struct {
	FirstName string `json:"first_name" binding:"required,lte=255"`
	Username  string `json:"username" binding:"required,lte=255"`
	Email     string `json:"email" binding:"required,email,lte=320"`
}

type UserFull struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	FirstName string    `json:"first_name"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
}

var UserColumns = [5]string{"id", "created_at", "first_name", "username", "email"}

func UserCreateNew(ctx context.Context, conn *pgx.Conn, newUser UserShort) (UserFull, error) {
	id := uuid.New()
	var createdUser UserFull
	err := conn.QueryRow(ctx,
		"INSERT INTO blog_user(id, first_name, username, email) VALUES($1, $2, $3, $4) RETURNING id, created_at, first_name, username, email",
		id, newUser.FirstName, newUser.Username, newUser.Email).Scan(
		&createdUser.ID,
		&createdUser.CreatedAt,
		&createdUser.FirstName,
		&createdUser.Username,
		&createdUser.Email)
	return createdUser, err
}

func UserGetSlice(ctx context.Context, conn *pgx.Conn, orderBy string) ([]UserFull, error) {
	usersSlice := make([]UserFull, 0)
	rows, err := conn.Query(ctx, fmt.Sprintf("SELECT * FROM blog_user ORDER BY %s ASC", orderBy))
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

func UserGetByID(ctx context.Context, conn *pgx.Conn, ID uuid.UUID) (UserFull, error) {
	var user UserFull
	err := conn.QueryRow(ctx, "SELECT * FROM blog_user WHERE id=($1)", ID).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.FirstName,
		&user.Username,
		&user.Email)
	return user, err
}

func UserUpdateByID(ctx context.Context, conn *pgx.Conn, ID uuid.UUID, newValues UserShort) (UserFull, error) {
	var user UserFull
	err := conn.QueryRow(ctx,
		"UPDATE blog_user SET first_name=($2), username=($3), email=($4) WHERE id=($1) RETURNING id, created_at, first_name, username, email",
		ID, newValues.FirstName, newValues.Username, newValues.Email).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.FirstName,
		&user.Username,
		&user.Email)
	return user, err
}

func UserDeleteByID(ctx context.Context, conn *pgx.Conn, ID uuid.UUID) error {
	result, err := conn.Exec(ctx, "DELETE FROM blog_user WHERE id=($1)", ID)
	if err != nil {
		return err
	} else {
		if result.RowsAffected() != 1 {
			return errors.New("No row found to delete")
		}
	}
	return nil
}
