package admin

import (
	"context"
	"delivery/configs"
	"delivery/constants"
	"delivery/entities"
	"delivery/logger"
	"delivery/storage"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type AdminController interface {
	CreateAPI(ctx context.Context, data *entities.API) error
	CreateAPIDetails(ctx context.Context, data *entities.APIDetails) error
	ExecuteAPI(ctx context.Context, apiName string, method string) (interface{}, error)
}

type adminController struct {
	log     logger.LoggerI
	storage storage.Storage
	cfg     *configs.Configuration
	redis   *redis.Client
}

func NewAdminController(log logger.LoggerI, storage storage.Storage, redis *redis.Client) AdminController {
	return &adminController{
		log:     log,
		storage: storage,
		cfg:     configs.Config(),
		redis:   redis,
	}
}

func (ac *adminController) CreateAPI(cxt context.Context, data *entities.API) error {
	return ac.storage.Admin().CreateAPI(cxt, data)
}

func (ac *adminController) CreateAPIDetails(cxt context.Context, data *entities.APIDetails) error {
	return ac.storage.Admin().CreateAPIDetails(cxt, data)
}

func (ac *adminController) ExecuteAPI(cxt context.Context, apiName string, method string) (interface{}, error) {
	res, err := ac.storage.Admin().GetAPI(cxt, apiName, method)
	if err != nil {
		return nil, err
	}
	result := make(map[string]interface{}) // Natijalarni saqlash uchun map

	for _, v := range res {
		var value interface{} // Query natijasini saqlash uchun

		switch v.QueryType {
		case constants.SELECT_LIST:
			var result []map[string]interface{}
			err := ac.storage.Admin().SelectList(cxt, v, nil, &result)
			if err != nil {
				fmt.Println("Error:", err)
				return nil, err
			}
			value = result
		case constants.SELECT_ONE:
			result := make(map[string]interface{})
			err := ac.storage.Admin().SelectOne(cxt, v, nil, &result)
			if err != nil {
				fmt.Println("Error:", err)
				return nil, err
			}
			value = result
		case constants.INSERT_ITEM:
			err := ac.storage.Admin().InsertItem(cxt, v)
			if err != nil {
				fmt.Println("Error:", err)
				return nil, err
			}
			value = "Insert success"
		case constants.UPDATE_ITEM:
			err := ac.storage.Admin().UpdateItem(cxt, v)
			if err != nil {
				fmt.Println("Error:", err)
				return nil, err
			}
			value = "Update success"
		}

		// Har bir natijani `v.name` kaliti bilan `result` map-ga qo‘shish
		result[v.Name] = value
	}

	return result, nil
}
