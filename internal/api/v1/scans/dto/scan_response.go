package dto

import "github.com/marrcoribeiro/security-scanner-api/internal/domain"

type ScanResponse struct {
	Id       string            `json:"id"`
	Path     string            `json:"path"`
	Findings []FindingResponse `json:"findings"`
	Done     bool              `json:"done"`
}

func ToResponse(scan *domain.Scan) *ScanResponse {
	response := &ScanResponse{
		Id:       scan.Id,
		Path:     scan.Path,
		Findings: ToFindingResponses(scan.Findings),
		Done:     scan.Done,
	}
	return response
}
