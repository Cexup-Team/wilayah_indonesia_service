package services

import (
	"api-test/helper"
	"api-test/model/web"
	"api-test/repository"
	"context"
	"database/sql"

	"github.com/go-playground/validator"
)

type RegionServiceImpl struct {
	RegionRepository repository.RegionRepository
	DB               *sql.DB
	validate         *validator.Validate
}

func NewRegionService(RegionRepository repository.RegionRepository, DB *sql.DB, validate *validator.Validate) RegionService {
	return &RegionServiceImpl{RegionRepository: RegionRepository, DB: DB, validate: validate}
}

func (c *RegionServiceImpl) FindAll(ctx context.Context) []web.RegionResponse {
	tx, err := c.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	regions := c.RegionRepository.FindAll(ctx, tx)

	return helper.ToRegionResponses(regions)
}

func (c *RegionServiceImpl) Find(ctx context.Context, id string) []web.RegionResponse {
	tx, err := c.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	regions := c.RegionRepository.Find(ctx, tx, id)

	return helper.ToRegionResponses(regions)
}
