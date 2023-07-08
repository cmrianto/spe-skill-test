package speSkillTest

import (
	"context"
	"speSkillTest/domain/request"
	"speSkillTest/domain/response"
)

type SpeSkillTestUsecaseInterface interface {
	Ping(ctx context.Context) (*response.Ping, error)
	NarcissisticNumber(ctx context.Context, req *request.NarcissisticNumberRequest) (*response.NarcissisticNumberResponse, error)
	ParityOutlier(ctx context.Context, req *request.ParityOutlierRequest) (*response.ParityOutlierResponse, error)
	NeedleInHaystack(ctx context.Context, req *request.NeedleInHaystackRequest) (*response.NeedleInHaystackResponse, error)
	BlueOcean(ctx context.Context, req *request.BlueOceanRequest) (*response.BlueOceanResponse, error)
}
