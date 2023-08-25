package model

import (
	"time"

	"github.com/lib/pq"
)

type Debug struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Title     string         `json:"title" gorm:"not null"`
	Body      string         `json:"body" gorm:"not null"`
	Links     pq.StringArray `json:"links" gorm:"type:text[]"`
	Techs     pq.StringArray `json:"techs" gorm:"type:text[]"`
	Cause     string         `json:"cause" gorm:"not null"`
	Resolve   string         `json:"resolve" gorm:"not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	User      User           `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	UserId    uint           `json:"user_id" gorm:"not null"`
}

type DebugResponse struct{
	ID        uint           `json:"id" gorm:"primaryKey"`
	Title     string         `json:"title" gorm:"not null"`
	Body      string         `json:"body" gorm:"not null"`
	Links     pq.StringArray `json:"links" gorm:"type:text[]"`
	Techs     pq.StringArray `json:"techs" gorm:"type:text[]"`
	Cause     string         `json:"cause" gorm:"not null"`
	Resolve   string         `json:"resolve" gorm:"not null"`
	User      User           `json:"user"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}