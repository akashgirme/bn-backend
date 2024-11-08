package sqlstore

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/akashgirme/bn-backend/internal/domain"
	"github.com/akashgirme/bn-backend/internal/model"
	"github.com/pkg/errors"
)

type UserRepository struct {
	db *sql.DB
}

func (r *UserRepository) GetByPhone(ctx context.Context, phone string) (*domain.User, error) {
	query := `
        SELECT id, phone_number, email, name, picture, created_at, updated_at 
        FROM users 
        WHERE phone_number = $1
    `

	var user domain.User
	err := r.db.QueryRowContext(ctx, query, phone).Scan(
		&user.ID,
		&user.PhoneNumber,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, repositories.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (us SqlUserStore) GetByEmail(email string) (*model.User, error) {
	query := us.usersQuery.Where("Email = lower(?)", email)

	queryString, args, err := query.ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "get_by_email_tosql")
	}

	user := model.User{}
	if err := us.GetReplicaX().Get(&user, queryString, args...); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.Wrap(us.NewErrNotFound("User", fmt.Sprintf("email=%s", email)), "failed to find User")
		}

		return nil, errors.Wrapf(err, "failed to get User with email=%s", email)
	}

	return &user, nil
}
