package usecase

import (
	"context"
	"math"
	"speSkillTest/domain/request"
	"speSkillTest/domain/response"
	"strconv"
)

func (s *service) NarcissisticNumber(ctx context.Context, req *request.NarcissisticNumberRequest) (*response.NarcissisticNumberResponse, error) {
	result := false

	digitRequest, _ := strconv.ParseFloat(req.Number, 64)

	pow := float64(len(req.Number))

	resultChecker := float64(0)

	for _, v := range req.Number {
		vFloat, _ := strconv.ParseFloat(string(v), 64)
		currentResult := math.Pow(vFloat, pow)
		resultChecker = resultChecker + currentResult
	}

	if digitRequest == resultChecker {
		result = true
	}

	return &response.NarcissisticNumberResponse{
		IsNarcissisticNumber: result,
	}, nil
}
