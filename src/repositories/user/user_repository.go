package userRepository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/kolaboradev/halo-suster/src/helper"
	userEntity "github.com/kolaboradev/halo-suster/src/models/entities/user"
	userRequest "github.com/kolaboradev/halo-suster/src/models/web/request/user"
)

type UserRepository struct{}

func NewUserRepository() UserRepositoryInterface {
	return &UserRepository{}
}

func (repository *UserRepository) SaveIt(ctx context.Context, tx *sql.Tx, user *userEntity.User) *userEntity.User {
	query := "INSERT INTO users (id, nip, name, password, role, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)"
	_, err := tx.ExecContext(ctx, query, user.Id, user.Nip, user.Name, user.Password.String, user.Role, user.CreatedAt, user.UpdatedAt)
	helper.ErrorIfPanic(err)
	return user
}

func (repository *UserRepository) GetUserItByNip(ctx context.Context, tx *sql.Tx, nip string) (userEntity.User, error) {
	query := "SELECT id, nip, name, password, role, created_at, updated_at FROM users WHERE nip = $1"

	rows, err := tx.QueryContext(ctx, query, nip)
	helper.ErrorIfPanic(err)
	defer rows.Close()

	user := userEntity.User{}
	if rows.Next() {
		err = rows.Scan(&user.Id, &user.Nip, &user.Name, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt)
		helper.ErrorIfPanic(err)
		return user, nil
	} else {
		return user, errors.New("user it not found")
	}
}

func (repository *UserRepository) SaveNurse(ctx context.Context, tx *sql.Tx, user *userEntity.User) *userEntity.User {
	query := "INSERT INTO users (id, nip, name, role, identity_card_scan_image, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)"

	_, err := tx.ExecContext(ctx, query, user.Id, user.Nip, user.Name, user.Role, user.IdentityCardScanImage.String, user.CreatedAt, user.UpdatedAt)
	helper.ErrorIfPanic(err)

	return user
}

func (repository *UserRepository) GetUserNurseByNip(ctx context.Context, tx *sql.Tx, nip string) (userEntity.User, error) {
	query := "SELECT id, nip, name, role, password, identity_card_scan_image, created_at, updated_at FROM users WHERE nip = $1"

	rows, err := tx.QueryContext(ctx, query, nip)
	helper.ErrorIfPanic(err)
	defer rows.Close()

	user := userEntity.User{}
	if rows.Next() {
		err = rows.Scan(&user.Id, &user.Nip, &user.Name, &user.Role, &user.Password, &user.IdentityCardScanImage, &user.CreatedAt, &user.UpdatedAt)
		helper.ErrorIfPanic(err)
		return user, nil
	} else {
		return user, errors.New("user nurse not found")
	}
}

func (repository *UserRepository) GetAllUser(ctx context.Context, tx *sql.Tx, filters userRequest.UserFilter) []userEntity.User {
	query := "SELECT id, nip, name, created_at, updated_at FROM users WHERE 1=1"

	var args []interface{}
	argIndex := 1

	if filters.UserId != "" {
		query += fmt.Sprintf(" AND id = $%d", argIndex)
		args = append(args, filters.UserId)
		argIndex++
	}
	if filters.Name != "" {
		fmt.Println(filters.Name)
		query += fmt.Sprintf(" AND name ILIKE '%%' || $%d || '%%'", argIndex)
		args = append(args, filters.Name)
		argIndex++
	}
	if filters.Nip != nil {
		query += fmt.Sprintf(" AND nip LIKE $%d || '%%'", argIndex)
		args = append(args, filters.Nip)
		argIndex++
	}
	if filters.Role != "" {
		if filters.Role == "it" {
			query += fmt.Sprintf(" AND role = $%d", argIndex)
			args = append(args, filters.Role)
			argIndex++
		}
		if filters.Role == "nurse" {
			query += fmt.Sprintf(" AND role = $%d", argIndex)
			args = append(args, filters.Role)
			argIndex++
		}
	}
	if filters.CreatedAt != "" {
		if filters.CreatedAt == "asc" {
			query += " ORDER BY created_at ASC"
		}
		if filters.CreatedAt == "desc" {
			query += " ORDER BY created_at DESC"
		}
	}

	if filters.Offset >= 0 {
		query += fmt.Sprintf(" OFFSET $%d", argIndex)
		args = append(args, filters.Offset)
		argIndex++
	}
	if filters.Limit >= 0 {
		if filters.Limit == 0 {
			filters.Limit = 5
		}
		query += fmt.Sprintf(" LIMIT $%d", argIndex)
		args = append(args, filters.Limit)
		argIndex++
	}

	fmt.Println(query)
	rows, err := tx.QueryContext(ctx, query, args...)
	helper.ErrorIfPanic(err)
	defer rows.Close()

	var users []userEntity.User
	for rows.Next() {
		user := userEntity.User{}
		rows.Scan(&user.Id, &user.Nip, &user.Name, &user.CreatedAt, &user.UpdatedAt)
		users = append(users, user)
	}
	return users
}

func (repository *UserRepository) EditNurseById(ctx context.Context, tx *sql.Tx, user userEntity.User) {
	query := "UPDATE users SET nip = $1, name = $2, updated_at = $3 WHERE id = $4"
	_, err := tx.ExecContext(ctx, query, user.Nip, user.Name, user.UpdatedAt, user.Id)
	helper.ErrorIfPanic(err)
}

func (repository *UserRepository) GetNurseByIdBool(ctx context.Context, tx *sql.Tx, id string) bool {
	query := "SELECT id FROM users WHERE id = $1"

	rows, err := tx.QueryContext(ctx, query, id)
	helper.ErrorIfPanic(err)
	defer rows.Close()

	if rows.Next() {
		return true
	} else {
		return false
	}

}

func (repository *UserRepository) DeleteNurseById(ctx context.Context, tx *sql.Tx, id string) {
	query := "DELETE FROM users where id = $1"
	_, err := tx.ExecContext(ctx, query, id)
	helper.ErrorIfPanic(err)
}

func (repository *UserRepository) GetNurseById(ctx context.Context, tx *sql.Tx, id string) (userEntity.User, error) {
	query := "SELECT id, nip FROM users WHERE id = $1"

	rows, err := tx.QueryContext(ctx, query, id)
	helper.ErrorIfPanic(err)
	defer rows.Close()

	user := userEntity.User{}
	if rows.Next() {
		err = rows.Scan(&user.Id, &user.Nip)
		helper.ErrorIfPanic(err)

		return user, nil
	} else {
		return user, errors.New("not found nurse")
	}
}

func (repository *UserRepository) AddNurseAccess(ctx context.Context, tx *sql.Tx, user userEntity.User) {
	query := "UPDATE users SET password = $1 WHERE id = $2"
	_, err := tx.ExecContext(ctx, query, user.Password.String, user.Id)
	helper.ErrorIfPanic(err)
}

func (repository *UserRepository) GetUserById(ctx context.Context, tx *sql.Tx, id string) (userEntity.User, error) {
	query := "SELECT id, nip, name FROM users WHERE id = $1"

	rows, err := tx.QueryContext(ctx, query, id)
	helper.ErrorIfPanic(err)
	defer rows.Close()

	user := userEntity.User{}
	if rows.Next() {
		err = rows.Scan(&user.Id, &user.Nip, &user.Name)
		helper.ErrorIfPanic(err)

		return user, nil
	} else {
		return user, errors.New("id not found")
	}
}
