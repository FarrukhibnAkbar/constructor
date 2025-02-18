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

func (r *adminRepo) SelectOne(ctx context.Context, req entities.APIDetails, arg []interface{}, dest *map[string]interface{}) error {
	rows, err := r.db.WithContext(ctx).Raw(req.Query, arg...).Rows()
	if err != nil {
		return utils.HandleDBError("SelectOne", err, r.db)
	}
	defer rows.Close()

	// Ustunlarni olish
	columns, err := rows.Columns()
	if err != nil {
		return utils.HandleDBError("SelectOne", err, r.db)
	}

	// Agar natija bo‘lsa, uni `map[string]interface{}` sifatida o‘qiymiz
	if rows.Next() {
		values := make([]interface{}, len(columns))
		valuePtrs := make([]interface{}, len(columns))

		for i := range columns {
			valuePtrs[i] = &values[i]
		}

		err := rows.Scan(valuePtrs...)
		if err != nil {
			return utils.HandleDBError("SelectOne", err, r.db)
		}

		result := make(map[string]interface{})
		for i, colName := range columns {
			result[colName] = values[i]
		}

		*dest = result
	}

	return nil
}


func (r *adminRepo) SelectList(ctx context.Context, req entities.APIDetails, arg []interface{}, dest *[]map[string]interface{}) error {
	rows, err := r.db.WithContext(ctx).Raw(req.Query, arg...).Rows()
	if err != nil {
		return utils.HandleDBError("SelectList", err, r.db)
	}
	defer rows.Close()

	// Ustunlarni olish
	columns, err := rows.Columns()
	if err != nil {
		return utils.HandleDBError("SelectList", err, r.db)
	}

	var results []map[string]interface{}

	// Har bir qator uchun natijani o'qish
	for rows.Next() {
		values := make([]interface{}, len(columns))
		valuePtrs := make([]interface{}, len(columns))

		for i := range columns {
			valuePtrs[i] = &values[i]
		}

		err := rows.Scan(valuePtrs...)
		if err != nil {
			return utils.HandleDBError("SelectList", err, r.db)
		}

		result := make(map[string]interface{})
		for i, colName := range columns {
			result[colName] = values[i]
		}

		// Natijani ro'yxatga qo'shish
		results = append(results, result)
	}

	*dest = results
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
