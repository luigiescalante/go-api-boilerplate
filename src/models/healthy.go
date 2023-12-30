package models

import (
	"go.api-boilerplate/config"
	"time"
)

type Healthy struct {
	Date   time.Time `json:"date"`
	Domain string    `json:"domain"`
}

func NewHealthy() *Healthy {
	cfg := config.GetConfig()
	return &Healthy{
		Date:   time.Now(),
		Domain: cfg.GetDomain(),
	}
}
