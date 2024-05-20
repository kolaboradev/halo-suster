package recordService

import (
	"context"

	recordRequest "github.com/kolaboradev/halo-suster/src/models/web/request/record"
	recordResponse "github.com/kolaboradev/halo-suster/src/models/web/response/record"
)

type RecordServiceInterface interface {
	Create(ctx context.Context, request recordRequest.Record)
	GetAll(ctx context.Context, filters recordRequest.RecordFilter) []recordResponse.RecordFilter
}
