package usecase

import (
	"context"
	"errors"
	"speSkillTest/domain/request"
	"speSkillTest/domain/response"
	"speSkillTest/lib/helper"

	"google.golang.org/grpc/codes"
)

func (s *service) NeedleInHaystack(ctx context.Context, req *request.NeedleInHaystackRequest) (*response.NeedleInHaystackResponse, error) {
	var (
		result   response.NeedleInHaystackResponse
		indexRes int64 = 0
		isFound  bool  = false
	)

	if req == nil {
		return &result, helper.Error(codes.InvalidArgument, "Invalid Argument:", errors.New("invalid argument"))
	}

	if len(req.Haystack) == 0 {
		return &result, helper.Error(codes.InvalidArgument, "Invalid Argument:", errors.New("haystack is empty"))
	}
	if len(req.Needle) == 0 {
		return &result, helper.Error(codes.InvalidArgument, "Invalid Argument:", errors.New("neelde is empty"))
	}

	for _, v := range req.Haystack {
		if req.Needle == v {
			isFound = true
			break
		}
		indexRes += 1
	}

	if !isFound {
		return &result, helper.Error(codes.NotFound, "Not Found:", errors.New("needle not found"))
	}

	return &response.NeedleInHaystackResponse{
		Index: indexRes,
	}, nil
}
