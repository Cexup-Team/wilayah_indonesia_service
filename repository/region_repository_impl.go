package repository

import (
	"api-test/helper"
	"api-test/model/domain"
	"context"
	"database/sql"
	"fmt"
)

type RegionRepositoryImpl struct {
}

func NewRegionRepository() RegionRepository {
	return &RegionRepositoryImpl{}
}

func (c *RegionRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Region {
	SQL := "SELECT * FROM wilayah_2020 WHERE CHAR_LENGTH(kode) < 3"
	row, err := tx.Query(SQL)
	helper.PanicIfError(err)

	var regions []domain.Region
	var order int = 0
	for row.Next() {
		order++
		region := domain.Region{}
		row.Scan(&region.Kode, &region.Name)
		helper.PanicIfError(err)

		regions = append(regions, region)
	}

	fmt.Println(order)
	return regions
}

func (c *RegionRepositoryImpl) Find(ctx context.Context, tx *sql.Tx, id string) []domain.Region {

	n := 2
	if len(id) < 8 {
		n = len(id) + 3
	} else {
		n = 13
	}

	SQL := "SELECT * FROM wilayah_2020 WHERE LEFT(kode, ? )= ? AND CHAR_LENGTH(kode)=? ORDER BY nama"
	row, err := tx.QueryContext(ctx, SQL, len(id), id, n)
	helper.PanicIfError(err)

	var regions []domain.Region
	var order int = 0
	for row.Next() {
		order++
		region := domain.Region{}
		row.Scan(&region.Kode, &region.Name)

		regions = append(regions, region)
	}

	fmt.Println(order)
	return regions
}
