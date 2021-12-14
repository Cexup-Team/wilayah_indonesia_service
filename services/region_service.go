package services

import (
	"api-test/model/web"
	"context"
)

type RegionService interface {
	// Create(ctx context.Context, request web.CategoryRequest) web.CategoryResponse
	// Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse
	// Delete(ctx context.Context, categoryId int)
	// FindId(ctx context.Context, categoryId int) web.CategoryResponse
	Find(ctx context.Context, id string) []web.RegionResponse
	FindAll(ctx context.Context) []web.RegionResponse
}
