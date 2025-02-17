package postgres

import (
	"context"
	"delivery/entities"
	"delivery/pkg/utils"
	"fmt"

	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

type adminRepo struct {
	db *gorm.DB
}

func NewAdmin(db *gorm.DB) *adminRepo {
	return &adminRepo{db: db}
}

func (r *adminRepo) CreateAPI(ctx context.Context, data *entities.API) error {
	return r.db.WithContext(ctx).Table("api").Create(data).Error
}

func (r *adminRepo) CreateAPIDetails(ctx context.Context, data *entities.APIDetails) error {
	return r.db.WithContext(ctx).Table("api_detail").Create(data).Error
}

func (r *adminRepo) GetAPI(ctx context.Context, apiName string, method string) ([]entities.APIDetails, error) {
	var details []entities.APIDetails
	var api entities.API
	err := r.db.WithContext(ctx).Table("api").Where("name = ? and method = ?", apiName, method).First(&api).Error
	if err != nil {
		return nil, utils.HandleDBError("GetAPI", err, r.db)
	}
	if api.ID == "" {
		return nil, fmt.Errorf("api not found")
	}
	err = r.db.WithContext(ctx).Table("api_detail").Where("api_id = ?", api.ID).Find(&details).Error
	if err != nil {
		return nil, utils.HandleDBError("GetAPI", err, r.db)
	}
	return details, nil
}

func (r *adminRepo) SelectOne(ctx context.Context, req entities.APIDetails, arg []interface{}, dest interface{}) error {
	var res []interface{}
	rows, err := r.db.WithContext(ctx).Raw(req.Query, arg...).Rows()
	if err != nil {
		return utils.HandleDBError("SelectOne", err, r.db)
	}
	err = rows.Scan(res...)
	if err != nil {
		return utils.HandleDBError("SelectOne", err, r.db)
	}
	fmt.Println(res...)
	return nil
}

func (r *adminRepo) SelectList(ctx context.Context, req entities.APIDetails, arg []interface{}, dest interface{}) error {
	err := r.db.WithContext(ctx).Raw(req.Query, arg...).Find(dest).Error
	if err != nil {
		return utils.HandleDBError("SelectList", err, r.db)
	}
	return nil
}

func (r *adminRepo) InsertItem(ctx context.Context, req entities.APIDetails) error {
	// err := r.db.WithContext(ctx).Table("api_detail").Create(&req).Error
	// if err != nil {
	// 	return utils.HandleDBError("InsertItem", err, r.db)
	// }
	return nil
}

func (r *adminRepo) UpdateItem(ctx context.Context, req entities.APIDetails) error {
	// err := r.db.WithContext(ctx).Table("api_detail").Where("id = ?", req.ID).Updates(&req).Error
	// if err != nil {
	// 	return utils.HandleDBError("UpdateItem", err, r.db)
	// }
	return nil
}
