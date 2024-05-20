package recordRepository

import (
	"context"
	"database/sql"

	recordEntity "github.com/kolaboradev/halo-suster/src/models/entities/record"
	recordRequest "github.com/kolaboradev/halo-suster/src/models/web/request/record"
)

type RecordRepositoryInterface interface {
	Save(ctx context.Context, tx *sql.Tx, record recordEntity.Record) recordEntity.Record
	GetAll(ctx context.Context, tx *sql.Tx, filters recordRequest.RecordFilter) []recordEntity.Record
}
