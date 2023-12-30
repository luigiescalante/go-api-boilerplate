package services

import "go.api-boilerplate/models"

func GetHealthy() *models.Healthy {
	hlt := models.NewHealthy()
	return hlt
}
