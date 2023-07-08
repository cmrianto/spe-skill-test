package usecase

import (
	"speSkillTest/config"

	external_interface "speSkillTest/lib/external_lib"

	speSkillTest "speSkillTest/domain/speSkillTest"

	"github.com/go-redis/redis/v8"
)

type service struct {
	env                string
	speSkillTestRepo   speSkillTest.SpeSkillTestRepoInterface
	redisClient        redis.UniversalClient
	config             *config.Config
	externalLibUsecase external_interface.ExternalLibInterface
}

type Dependencies struct {
	Env                string
	SpeSkillTestRepo   speSkillTest.SpeSkillTestRepoInterface
	RedisClient        redis.UniversalClient
	Config             *config.Config
	ExternalLibUsecase external_interface.ExternalLibInterface
}

func NewService(deps Dependencies) speSkillTest.SpeSkillTestUsecaseInterface {
	return &service{
		env:                deps.Env,
		speSkillTestRepo:   deps.SpeSkillTestRepo,
		redisClient:        deps.RedisClient,
		config:             deps.Config,
		externalLibUsecase: deps.ExternalLibUsecase,
	}
}
