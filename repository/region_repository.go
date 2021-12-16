package repository

import (
	"context"
	"database/sql"
	"wilayah_indonesia_service/model/domain"
)

type RegionRepository interface {
	// Save(ctx context.Context, tx *sql.Tx, category domain.Region) domain.Region
	// Update(ctx context.Context, tx *sql.Tx, category domain.Region) domain.Region
	// Delete(ctx context.Context, tx *sql.Tx, category domain.Region)
	// FindId(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Region, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Region
	Find(ctx context.Context, tx *sql.Tx, id string) []domain.Region
}
