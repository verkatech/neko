package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/verkatech/neko/pkg/config"
	"github.com/verkatech/neko/pkg/model"
	"github.com/verkatech/neko/pkg/router"
	"gopkg.in/urfave/cli.v2"
)

// Server is the cli command that runs our main web server
func Server() *cli.Command {
	return &cli.Command{
		Name:  "server",
		Usage: "Starts the neko web server",
		Before: func(c *cli.Context) error {
			err := model.SetupMongo()
			if err != nil {
				log.Info("Error connecting to mongodb", err)
				return err
			}
			return nil
		},
		Flags: []cli.Flag{

		},
		Action: func(c *cli.Context) error {
			ginEngine := gin.Default()
			router.InitRoutes(ginEngine)
			err := ginEngine.Run(fmt.Sprintf("%s:%d", config.Server.ServerHost, config.Server.ServerPort))
			if err != nil {
				log.Fatal("Failed running gin web server", err)
			}
			return nil
		},
	}
}
