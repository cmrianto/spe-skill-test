package cmd

import (
	"fmt"
	"log"
	"speSkillTest/config"
	"speSkillTest/lib/pkg/server"
	"speSkillTest/route"

	"speSkillTest/application"
	lib_tracer "speSkillTest/lib/pkg/tracer"

	"github.com/urfave/cli"
)

func runCommand(cfg *config.Config) func(*cli.Context) error {
	return func(c *cli.Context) error {
		t, fnCloser, err := lib_tracer.NewTracer()
		if err != nil {
			return err
		}
		defer fnCloser()
		srv := server.NewServer()
		srv.Use(lib_tracer.GinOpentracingMiddleware(t))
		app, err := application.Setup(cfg, c)
		if err != nil {
			return err
		}
		route.SetupRouter(srv, cfg, app)

		log.Println("Starting server " + cfg.Application.ServerPort)
		if err := srv.Run(fmt.Sprintf("%s:%s", cfg.Application.ServerHost, cfg.Application.ServerPort)); err != nil {
			return err
		}
		return nil
	}
}

func Cli(cfg *config.Config) *cli.App {
	clientApp := cli.NewApp()
	clientApp.Name = cfg.Application.ServiceName
	clientApp.Version = cfg.Application.ServiceVersion
	clientApp.Action = runCommand(cfg)
	return clientApp
}
