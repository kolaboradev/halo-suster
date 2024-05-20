package recordRepository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/kolaboradev/halo-suster/src/helper"
	recordEntity "github.com/kolaboradev/halo-suster/src/models/entities/record"
	recordRequest "github.com/kolaboradev/halo-suster/src/models/web/request/record"
)

type RecordRepository struct{}

func NewRecordRepository() RecordRepositoryInterface {
	return &RecordRepository{}
}

func (repository *RecordRepository) Save(ctx context.Context, tx *sql.Tx, record recordEntity.Record) recordEntity.Record {
	query := "INSERT INTO records (id, identity_number, symptoms, medications, created_by, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)"

	_, err := tx.ExecContext(ctx, query, record.Id, record.IdentityNumber, record.Symptoms, record.Medications, record.CreatedBy, record.CreatedAt, record.UpdatedAt)
	helper.ErrorIfPanic(err)

	return record
}

func (repository *RecordRepository) GetAll(ctx context.Context, tx *sql.Tx, filters recordRequest.RecordFilter) []recordEntity.Record {
	query := fmt.Sprint("SELECT r.identity_number, r.symptoms, r.medications, r.created_by, r.created_at " +
		"FROM records AS r JOIN patients AS p ON r.identity_number = p.identity_number " +
		"JOIN users AS u ON r.created_by = u.id WHERE 1=1")

	var args []interface{}
	argIndex := 1

	if filters.IdentityNumber != nil {
		query += fmt.Sprintf(" AND p.identity_number = $%d", argIndex)
		args = append(args, filters.IdentityNumber)
		argIndex++
	}

	if filters.UserId != "" {
		query += fmt.Sprintf(" AND u.id = $%d", argIndex)
		args = append(args, filters.UserId)
		argIndex++
	}

	if filters.Nip != "" {
		query += fmt.Sprintf(" AND u.nip = $%d", argIndex)
		args = append(args, filters.Nip)
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

	var records []recordEntity.Record
	for rows.Next() {
		record := recordEntity.Record{}
		rows.Scan(&record.IdentityNumber, &record.Symptoms, &record.Medications, &record.CreatedBy, &record.CreatedAt)
		records = append(records, record)
	}
	return records
}
