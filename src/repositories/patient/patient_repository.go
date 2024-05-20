package patientRepository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/kolaboradev/halo-suster/src/helper"
	patientEntity "github.com/kolaboradev/halo-suster/src/models/entities/patient"
	patientRequest "github.com/kolaboradev/halo-suster/src/models/web/request/patient"
)

type PatientRepository struct{}

func NewPatientRepo() PatientRepositoryInterface {
	return &PatientRepository{}
}

func (repository *PatientRepository) Save(ctx context.Context, tx *sql.Tx, patient *patientEntity.Patient) *patientEntity.Patient {
	query := "INSERT INTO patients (id, identity_number, phone_number, name, birth_date, gender, identity_card_scan_image, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)"
	_, err := tx.ExecContext(ctx, query, patient.Id, patient.IdentityNumber, patient.PhoneNumber, patient.Name, patient.BirthDate, patient.Gender, patient.IdentityCardScanImage, patient.CreatedAt, patient.UpdatedAt)
	helper.ErrorIfPanic(err)

	return patient
}

func (repository *PatientRepository) GetByIdentityNumberBool(ctx context.Context, tx *sql.Tx, identityNumber int64) bool {
	query := "SELECT identity_number FROM patients WHERE identity_number = $1"
	rows, err := tx.QueryContext(ctx, query, identityNumber)
	helper.ErrorIfPanic(err)
	defer rows.Close()

	if rows.Next() {
		return true
	} else {
		return false
	}
}

func (repository *PatientRepository) GetAll(ctx context.Context, tx *sql.Tx, filters patientRequest.PatientFilter) []patientEntity.Patient {
	query := "SELECT identity_number, phone_number, name, birth_date, gender, created_at, updated_at FROM patients WHERE 1=1"
	var args []interface{}
	argIndex := 1

	if filters.IdentityNumber != nil {
		query += fmt.Sprintf(" AND identity_number = $%d", argIndex)
		args = append(args, filters.IdentityNumber)
		argIndex++
	}
	if filters.Name != "" {
		query += fmt.Sprintf(" AND name ILIKE '%%' || $%d || '%%'", argIndex)
		args = append(args, filters.Name)
		argIndex++
	}
	if filters.PhoneNumber != "" {
		phoneNumber := "+" + filters.PhoneNumber
		fmt.Println(phoneNumber)
		query += fmt.Sprintf(" AND phone_number LIKE $%d || '%%'", argIndex)
		args = append(args, phoneNumber)
		argIndex++
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

	var users []patientEntity.Patient
	for rows.Next() {
		user := patientEntity.Patient{}
		err = rows.Scan(&user.IdentityNumber, &user.PhoneNumber, &user.Name, &user.BirthDate, &user.Gender, &user.CreatedAt, &user.UpdatedAt)
		helper.ErrorIfPanic(err)

		users = append(users, user)
	}
	return users
}

func (repository *PatientRepository) GetByIdentityNumber(ctx context.Context, tx *sql.Tx, id int64) (patientEntity.Patient, error) {
	query := "SELECT identity_number, phone_number, name, birth_date, gender, identity_card_scan_image FROM patients WHERE identity_number = $1"

	rows, err := tx.QueryContext(ctx, query, id)
	helper.ErrorIfPanic(err)
	defer rows.Close()

	patient := patientEntity.Patient{}
	if rows.Next() {
		err = rows.Scan(&patient.IdentityNumber, &patient.PhoneNumber, &patient.Name, &patient.BirthDate, &patient.Gender, &patient.IdentityCardScanImage)
		helper.ErrorIfPanic(err)
		return patient, nil
	} else {
		return patient, errors.New("identity number not found")
	}
}
