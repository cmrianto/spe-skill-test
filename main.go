package main

import (
	"log"
	"os"
	"speSkillTest/cmd"

	"speSkillTest/config"
)

func main() {
	cfg, err := config.Setup()
	if err != nil {
		log.Fatal("Cannot load config ", err.Error())
	}

	if cmd.Cli(cfg).Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
