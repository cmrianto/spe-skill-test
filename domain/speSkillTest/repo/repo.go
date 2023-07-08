package repo

import (
	"speSkillTest/domain/speSkillTest"

	"gorm.io/gorm"
)

type speSkillTestRepo struct {
	db *gorm.DB
}

func NewMysqlRepo(
	db *gorm.DB,
) speSkillTest.SpeSkillTestRepoInterface {
	return &speSkillTestRepo{
		db: db,
	}
}
