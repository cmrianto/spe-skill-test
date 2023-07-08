package route

import (
	"context"
	"net/http"
	"speSkillTest/application"
	"speSkillTest/config"
	"speSkillTest/domain/handler/http_handler"
	"speSkillTest/domain/speSkillTest/repo"
	"speSkillTest/domain/speSkillTest/usecase"
	"speSkillTest/lib/helper/http_response"
	"strings"

	"os"

	external_usecase "speSkillTest/lib/external_lib/usecase"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type healthCheckResponse struct {
	Data string `json:"data"`
}

func SetupRouter(r *gin.Engine, cfg *config.Config, app *application.Application) {
	r.Use(requestid.New())
	healthRes := healthCheckResponse{
		Data: "pong",
	}
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, healthRes)
	})
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, healthRes)
	})
	ctx := context.Background()

	skipAuthEnv := os.Getenv("SKIP_AUTH_ENV")
	skipAuthEnv = strings.ReplaceAll(skipAuthEnv, " ", "")

	// Repo
	db := &gorm.DB{}
	mysqlRepo := repo.NewMysqlRepo(
		db,
	)

	// ExternalLibInterface
	externalLibInterface := external_usecase.NewUsecase(external_usecase.Dependencies{
		Env:         cfg.Application.Env,
		Db:          db,
		Config:      cfg,
		RedisClient: app.RedisClient,
	})

	// Usecase
	speSkillTestUsecase := usecase.NewService(usecase.Dependencies{
		Env:                cfg.Application.Env,
		SpeSkillTestRepo:   mysqlRepo,
		RedisClient:        app.RedisClient,
		Config:             cfg,
		ExternalLibUsecase: externalLibInterface,
	})

	// Handler
	speSkillTestHandler := http_handler.NewSpeSkillTestHandler(ctx, speSkillTestUsecase)

	// APIs
	r.GET("/ping", ModeWrapper(speSkillTestHandler.Ping))
	r.GET("/narcissistic-number", ModeWrapper(speSkillTestHandler.NarcissisticNumber))
	r.GET("/parity-outlier", ModeWrapper(speSkillTestHandler.ParityOutlier))
	r.POST("/needle-in-haystack", ModeWrapper(speSkillTestHandler.NeedleInHaystack))
	r.POST("/blue-ocean", ModeWrapper(speSkillTestHandler.BlueOcean))
}

func AuthenticationSkipper(env string, skipAuthEnv []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var isSkipAuth bool
		for _, skipEnv := range skipAuthEnv {
			if env == skipEnv {
				isSkipAuth = true
				break
			}
		}

		if !isSkipAuth {
			c.Next()
			return
		}

		skipAuthHeader := c.GetHeader("X-SKIP-AUTH")
		if skipAuthHeader == "" {
			c.Next()
			return
		}

		skipAuthTokenHash := os.Getenv("SKIP_AUTH_TOKEN_HASH")
		if err := bcrypt.CompareHashAndPassword([]byte(skipAuthTokenHash), []byte(skipAuthHeader)); err != nil {
			c.Next()
			return
		}

		skipAuthSsoId := os.Getenv("SKIP_AUTH_SSO_ID")
		skipAuthUserRoleId := os.Getenv("SKIP_AUTH_USER_ROLE_ID")
		skipAuthRoleId := os.Getenv("SKIP_AUTH_ROLE_ID")
		skipAuthRoleName := os.Getenv("SKIP_AUTH_ROLE_NAME")
		skipAuthUserName := os.Getenv("SKIP_AUTH_USER_NAME")

		c.Set("skip-auth", struct{}{})
		c.Set("skip-auth-sso-id", skipAuthSsoId)
		c.Set("skip-auth-user-name", skipAuthUserName)
		c.Set("skip-auth-user-role-id", skipAuthUserRoleId)
		c.Set("skip-auth-role-id", skipAuthRoleId)
		c.Set("skip-auth-role-name", skipAuthRoleName)
		c.Next()
	}
}

func ModeWrapper(h gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if val, ok := c.Get("run-mode"); ok {
			mode, ok := val.(string)
			if ok {
				_, exists := c.Get("public-mode")
				if mode == "public" {
					if exists {
						h(c)
					} else {
						http_response.AbortError(status.Error(codes.PermissionDenied, "api access forbidden"), c)
					}
				} else {
					h(c)
				}
			}
		} else {
			h(c)
		}
	}
}
