package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
	"web_test01/types"
	app_err "web_test01/utility/errors"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type UserStore struct {
	logger *zap.SugaredLogger
	db     *sqlx.DB
}

type User struct {
	ID        int64     `json:"userId" db:"user_id"`
	FirstName string    `json:"firstName" db:"first_name"`
	LastName  string    `json:"lastName" db:"last_name"`
	Email     string    `json:"userEmail" db:"email_id"`
	Password  string    `json:"-" db:"password"`
	CreatedAt time.Time `json:"-" db:"created_at"`
	UpdatedAt time.Time `json:"-" db:"updated_at"`
}

func (us *UserStore) Create(payload types.RegisterPayload) (int64, error) {
	query := `
		INSERT INTO users(first_name, last_name, email_id, password) 
		VALUES($1, $2, $3, $4)
	`
	resp, err := us.db.Exec(query, payload.FirstName, payload.LastName, payload.Email, payload.Password)
	if err != nil {
		return 0, nil
	}

	rowsAffected, err := resp.RowsAffected()
	if err != nil {
		return 0, nil
	}

	fmt.Println(rowsAffected)
	return rowsAffected, nil
}

func (us *UserStore) GetUserByEmail(email string) (*User, error) {
	var user User

	query := `
		SELECT user_id, first_name, last_name, email_id, password FROM users where email_id = $1
	`
	if err := us.db.Get(&user, query, email); err != nil {
		us.logger.Errorw("store GetUserByEmail error", "err", err)
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, app_err.ErrNoRecords
		default:
			return nil, err
		}
	}

	return &user, nil
}

func (us *UserStore) GetUserById(id int64) (*User, error) {
	var user User

	query := `
		SELECT user_id, first_name, last_name, email_id, password FROM users where user_id = $1
	`
	if err := us.db.Get(&user, query, id); err != nil {
		us.logger.Errorw("store GetUserById error", "err", err)
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, app_err.ErrNoRecords
		default:
			return nil, err
		}
	}

	return &user, nil
}
