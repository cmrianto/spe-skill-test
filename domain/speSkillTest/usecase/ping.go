package usecase

import (
	"context"
	"speSkillTest/domain/response"
)

func (s *service) Ping(ctx context.Context) (*response.Ping, error) {
	return &response.Ping{
		Message: "pong",
	}, nil
}
