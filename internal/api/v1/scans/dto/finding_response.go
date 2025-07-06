package dto

import "github.com/marrcoribeiro/security-scanner-api/internal/domain"

type FindingResponse struct {
	// Rule is the identifier for the rule that was violated.
	Rule    string `json:"rule"`
	// File is the path to the file where the finding was detected.
	File    string `json:"file"`
	// Message is a description of the finding.
	Message string `json:"message"`
	// Line is the line number in the file where the finding was detected.
	Line    int    `json:"line"`
}

func ToFindingResponse(f domain.Finding) FindingResponse {
	return FindingResponse{
		Line:    f.Line,
		Rule:    f.Rule,
		File:    f.File,
		Message: f.Message,
	}
}

func ToFindingResponses(list []domain.Finding) []FindingResponse {
	res := make([]FindingResponse, 0, len(list))
	for _, f := range list {
		res = append(res, ToFindingResponse(f))
	}
	return res
}