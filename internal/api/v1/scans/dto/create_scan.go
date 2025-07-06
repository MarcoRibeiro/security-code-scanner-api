package dto

import (
	"github.com/marrcoribeiro/security-scanner-api/internal/domain"
)

type CreateScanRequest struct {
	Path         string   `json:"path" binding:"required"` 
	Configuration *Configuration `json:"configuration" binding:"required"`
}

func (req *CreateScanRequest) ToDomain() *domain.Scan {
	scan := &domain.Scan{
		Path:    req.Path,
	}

	if req.Configuration != nil {
		scan.Configuration = &domain.Configuration{
			Exclude: req.Configuration.Exclude,
		}
	}

	return scan
}