package dto

import "github.com/marrcoribeiro/security-scanner-api/internal/domain"

type Configuration struct {
	Exclude []string `json:"exclude"`
}

func (c *Configuration) ToDomain() *domain.Configuration {
	return &domain.Configuration{
		Exclude: c.Exclude,
	}
}
