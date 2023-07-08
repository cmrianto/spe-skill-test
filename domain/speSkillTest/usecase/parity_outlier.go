package usecase

import (
	"context"
	"fmt"
	"speSkillTest/domain/request"
	"speSkillTest/domain/response"
	"strconv"
)

func (s *service) ParityOutlier(ctx context.Context, req *request.ParityOutlierRequest) (*response.ParityOutlierResponse, error) {
	var (
		result      string
		evenNumbers []int64
		oddNumbers  []int64
	)

	for _, v := range req.Numbers {
		tempNumber, _ := strconv.Atoi(v)
		if tempNumber%2 == 0 {
			evenNumbers = append(evenNumbers, int64(tempNumber))
		} else {
			oddNumbers = append(oddNumbers, int64(tempNumber))
		}
	}

	if len(evenNumbers) > len(oddNumbers) {
		for _, v := range oddNumbers {
			if len(result) == 0 {
				result = fmt.Sprintf("%d", v)
			} else {
				result += fmt.Sprintf(", %d", v)
			}
		}
	} else if len(oddNumbers) > len(evenNumbers) {
		for _, v := range evenNumbers {
			if len(result) == 0 {
				result = fmt.Sprintf("%d", v)
			} else {
				result += fmt.Sprintf(", %d", v)
			}
		}
	} else if len(oddNumbers) == 0 {
		result = "false"
	} else if len(evenNumbers) == 0 {
		result = "false"
	} else if len(evenNumbers) == 0 && len(oddNumbers) == 0 {
		result = "false"
	} else if len(evenNumbers) == len(oddNumbers) {
		result = "false"
	}

	return &response.ParityOutlierResponse{
		Outlier: result,
	}, nil
}
