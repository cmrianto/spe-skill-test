package http_handler

import (
	"context"

	speSkillTest "speSkillTest/domain/speSkillTest"

	"github.com/gin-gonic/gin"
)

type handler struct {
	ctx     context.Context
	usecase speSkillTest.SpeSkillTestUsecaseInterface
}

type SpeSkillTestHandler interface {
	// Service
	Ping(*gin.Context)
	NarcissisticNumber(*gin.Context)
	ParityOutlier(*gin.Context)
	NeedleInHaystack(*gin.Context)
	BlueOcean(*gin.Context)
}

func NewSpeSkillTestHandler(
	ctx context.Context,
	usecase speSkillTest.SpeSkillTestUsecaseInterface,
) SpeSkillTestHandler {
	return &handler{
		ctx:     ctx,
		usecase: usecase,
	}
}
