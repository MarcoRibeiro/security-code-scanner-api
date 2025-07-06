package dto

import "github.com/marrcoribeiro/security-scanner-api/internal/domain"

type FindingResponse struct {
	Rule    string `json:"rule"`
	File    string `json:"file"`
	Message string `json:"message"`
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