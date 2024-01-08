package models

import (
	"go.api-boilerplate/config"
	"time"
)

type Healthy struct {
	APP    string    `json:"app"`
	Date   time.Time `json:"date"`
	Domain string    `json:"domain"`
}

func NewHealthy() *Healthy {
	cfg := config.GetConfig()
	return &Healthy{
		APP:    "Api boilerplate",
		Date:   time.Now(),
		Domain: cfg.GetDomain(),
	}
}
