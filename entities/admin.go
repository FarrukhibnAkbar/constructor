package entities

type API struct {
	ID     string `json:"id" gorm:"column:id"`
	Name   string `json:"name" gorm:"column:name"`
	Method string `json:"method" gorm:"column:method"`
	// CreatedBy sql.NullString `gorm:"column:created_by"`
	// UpdatedBy sql.NullString `gorm:"column:updated_by"`
}

type APIDetails struct {
	ID        string `json:"id" gorm:"column:id"`
	APIID     string `json:"api_id" gorm:"column:api_id"`
	Name      string `json:"name" gorm:"column:name"`
	QueryType string `json:"query_type" gorm:"column:query_type"`
	Query     string `json:"query" gorm:"column:query"`
	// CreatedBy sql.NullString `gorm:"column:created_by"`
	// UpdatedBy sql.NullString `gorm:"column:updated_by"`
}

type APIItems struct {
	Details []APIDetails
	Count   int
}
