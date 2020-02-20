package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/verkatech/neko/pkg/model"
	"gopkg.in/urfave/cli.v2"
)

// CreateAdminUser is the cli command that creates a new admin user
func NewAdminUser() *cli.Command {
	return &cli.Command{
		Name:  "new_admin",
		Usage: "Creates a new admin user",
		Before: func(c *cli.Context) error {
			err := model.SetupMongo()
			if err != nil {
				log.Fatal("Error connecting to mongodb", err)
				return err
			}
			return nil
		},
		Flags: []cli.Flag{

		},
		Action: func(c *cli.Context) error {
			fmt.Print("Not Implemented yet")
			return nil
		},
	}
}
