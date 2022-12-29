package main

// BusinessMetric struct
type BusinessMetric struct {
	ID                uint     `json:"id" form:"id" query:"id"`
	Name              string   `json:"name" form:"name" query:"name"`
	Description       string   `json:"description" form:"description" query:"description"`
	Version           uint     `json:"version" form:"version" query:"version" gorm:"default:1"`
	PreviousVersionID uint     `json:"previous_version_id" form:"previous_version_id" query:"previous_version_id"`
	LatestVersion     bool     `json:"-" form:"-" query:"latest_version" gorm:"default:true"`
	CreatedAt         int      `json:"-" form:"created_at" query:"created_at"`
	UpdatedAt         int      `json:"-" form:"updated_at" query:"updated_at"`
	DeletedAt         int      `json:"-" form:"deleted_at" query:"deleted_at"`
	Events            []*Event `json:"events" form:"events" query:"events" gorm:"many2many:business_metrics_events;"`
}

// Event struct
type Event struct {
	ID                uint   `json:"id" form:"id" query:"id"`
	Name              string `json:"name" form:"name" query:"name"`
	Description       string `json:"description" form:"description" query:"description"`
	Type              string `json:"type" form:"type" query:"type"`
	Required          bool   `json:"required" form:"required" query:"required" gorm:"default:false"`
	Version           uint   `json:"version" form:"version" query:"version" gorm:"default:1"`
	PreviousVersionID uint   `json:"previous_version_id" form:"previous_version_id" query:"previous_version_id"`
	LatestVersion     bool   `json:"-" form:"latest_version" query:"latest_version" gorm:"default:true"`
	StandardEvent     bool   `json:"standard_event" form:"standard_event" query:"standard_event" gorm:"default:false"`
	CreatedAt         int    `json:"-" form:"created_at" query:"created_at"`
	UpdatedAt         int    `json:"-" form:"updated_at" query:"updated_at"`
	DeletedAt         int    `json:"-" form:"deleted_at" query:"deleted_at"`
}
