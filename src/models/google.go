package models

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
)

type GoogleIDTokenShort struct {
	Sub string `json:"sub"`
	Iat int64  `json:"iat"`
	Exp int64  `json:"exp"`
}

type GoogleIDTokenFull struct {
	Sub  string    `json:"sub"`
	User uuid.UUID `json:"user"`
	Iat  int64     `json:"iat"`
	Exp  int64     `json:"exp"`
}

func GoogleIDTokenCreateOrUpdate(ctx context.Context, conn *pgx.Conn, newToken GoogleIDTokenShort) error {
	_, err := conn.Exec(ctx,
		"INSERT INTO google_id_token(sub, iat, exp) VALUES($1, $2, $3) ON CONFLICT (sub) DO UPDATE SET iat=($2), exp=($3)",
		newToken.Sub, newToken.Iat, newToken.Exp)
	return err
}
