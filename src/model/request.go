package model

import "gorm.io/gorm"

type Request struct {
	gorm.Model
	IP        string  `gorm:"size:64;"`
	URI       string  `gorm:"size:255"`
	City      string  `gorm:"size:255"`
	Provinces string  `gorm:"size:255"`
	Country   string  `gorm:"size:255"`
	Continent string  `gorm:"size:255"`
	TimeZone  string  `gorm:"size:64;"`
	Latitude  float64 `gorm:"type:float;"`
	Longitude float64 `gorm:"type:float;"`
}
