package repo

import (
	"context"
	"delivery/entities"
)

// IAdminStorage account storage interface
type IAdminStorage interface {
	CreateAPI(ctx context.Context, data *entities.API) error
	CreateAPIDetails(ctx context.Context, data *entities.APIDetails) error
	GetAPI(ctx context.Context, apiName string, method string) ([]entities.APIDetails, error)
	SelectOne(ctx context.Context, req entities.APIDetails, arg []interface{}, dest *map[string]interface{}) error
	SelectList(ctx context.Context, req entities.APIDetails, arg []interface{}, dest *[]map[string]interface{}) error
	InsertItem(ctx context.Context, req entities.APIDetails) error
	UpdateItem(ctx context.Context, req entities.APIDetails) error
}
