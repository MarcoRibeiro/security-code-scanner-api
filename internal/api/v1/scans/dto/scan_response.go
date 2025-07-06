package dto

import "github.com/marrcoribeiro/security-scanner-api/internal/domain"

type ScanResponse struct {
	// Id is the unique identifier for the scan.
	Id       string            `json:"id"`
	// Path is the file or directory that was scanned.
	Path     string            `json:"path"`
	// Findings contains the results of the scan.
	Findings []FindingResponse `json:"findings"`
	// Done indicates whether the scan has completed.
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
