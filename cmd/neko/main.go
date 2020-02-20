package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/verkatech/neko/pkg/config"
	"gopkg.in/urfave/cli.v2"
	"os"
	"strings"
	"time"
)

// setupLogging sets logging level for logrus
func setupLogging() {
	switch strings.ToLower(config.Logging.Level) {
	case "debug":
		log.SetLevel(log.DebugLevel)
	default:
		log.SetLevel(log.InfoLevel)
	}
}

// cliFlags returns global cli flags
func cliFlags() []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:        "debug",
			Value:       true,
			Usage:       "Activate debug information",
			EnvVars:     []string{"NEKO_DEBUG"},
			Destination: &config.Server.Debug,
		},
		&cli.StringFlag{
			Name:        "logging-level",
			Value:       "info",
			Usage:       "set logging level",
			EnvVars:     []string{"NEKO_LOG_LEVEL"},
			Destination: &config.Logging.Level,
		},
		&cli.StringFlag{
			Name:        "mongo-uri",
			Value:       "mongodb://localhost:27019",
			Usage:       "Mongo database uri",
			EnvVars:     []string{"NEKO_MONGO_URI"},
			Destination: &config.Database.Uri,
		},
		&cli.StringFlag{
			Name:        "mongo-database",
			Value:       "neko",
			Usage:       "Mongo database name",
			EnvVars:     []string{"NEKO_MONGO_DATABASE"},
			Destination: &config.Database.DatabaseName,
		},
		&cli.StringFlag{
			Name:        "static-path",
			Value:       "./client/build",
			Usage:       "Static files path",
			EnvVars:     []string{"NEKO_STATIC_PATH"},
			Destination: &config.Server.StaticPath,
		},
		&cli.IntFlag{
			Name:        "server-port",
			Value:       8062,
			Usage:       "Web server port",
			EnvVars:     []string{"NEKO_SERVER_PORT", "PORT"},
			Destination: &config.Server.ServerPort,
		},
		&cli.StringFlag{
			Name:        "server-host",
			Value:       "0.0.0.0",
			Usage:       "Web server host",
			EnvVars:     []string{"NEKO_SERVER_HOST"},
			Destination: &config.Server.ServerHost,
		},
		&cli.IntFlag{
			Name:        "token-ttl",
			Value:       24 * 8,
			Usage:       "Auth token time to live in hours",
			EnvVars:     []string{"NEKO_TOKEN_TTL"},
			Destination: &config.Server.TokenTimeToLive,
		},
	}
}

func main() {
	app := &cli.App{
		Name:     "Neko",
		Usage:    "neko web server",
		Compiled: time.Now(),
		Version:  "0.3",
		Authors: []*cli.Author{
			{
				Name:  "Iman Daneshi",
				Email: "emandaneshikohan@gmail.com",
			},
		},
		Flags: cliFlags(),
		Commands: []*cli.Command{
			Server(),
			NewAdminUser(),
		},
		Before: func(c *cli.Context) error {
			setupLogging()
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal("failed starting the web server")
	}
}
