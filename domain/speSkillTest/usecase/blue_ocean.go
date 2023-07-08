package usecase

import (
	"context"
	"errors"
	"speSkillTest/domain/request"
	"speSkillTest/domain/response"
	"speSkillTest/lib/helper"

	"google.golang.org/grpc/codes"
)

func (s *service) BlueOcean(ctx context.Context, req *request.BlueOceanRequest) (*response.BlueOceanResponse, error) {
	var (
		result response.BlueOceanResponse
	)

	if req == nil {
		return &result, helper.Error(codes.InvalidArgument, "Invalid Argument:", errors.New("invalid argument"))
	}

	if len(req.BlueOcean) == 0 {
		return &result, helper.Error(codes.InvalidArgument, "Invalid Argument:", errors.New("blue ocean is empty"))
	}
	if len(req.Remove) == 0 {
		result.BlueOcean = req.BlueOcean
		return &result, nil
	}

	removeMarker := make(map[int64]struct{}, len(req.Remove))
	for _, v := range req.Remove {
		removeMarker[v] = struct{}{}
	}

	for _, blueOcean := range req.BlueOcean {
		if _, found := removeMarker[blueOcean]; !found {
			result.BlueOcean = append(result.BlueOcean, blueOcean)
		}
	}

	return &result, nil
}
