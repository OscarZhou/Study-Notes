package models

import "time"

type CurrencyRelation struct {
	FEICurrencyCode  string `gorm:"primary_key"`
	TRKDCurrencyCode string `gorm:"primary_key"`
	TRKDRIC          string
	CurrencyName     string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        *time.Time `json:"-" sql:"index"`
}
