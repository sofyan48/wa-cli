package cmd

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/orn/wa-cli/src/libs/waapi"
	"github.com/urfave/cli"
)

// CLIMapping ...
type CLIMapping struct {
	chat waapi.WhatsappWebClient
}

// CLIMappingHandler ...
func CLIMappingHandler() *CLIMapping {
	return &CLIMapping{
		chat: waapi.NewWhatsappWebClient(),
	}
}

// ArgsMapping object mapping
type ArgsMapping struct {
	EnvPath string
	// login params
	QrPath string
}

// Args Glabal Acces args command
var Args ArgsMapping
var app *cli.App

// AppCommands All Command line app
func AppCommands() *cli.App {
	app := Init()
	handler := CLIMappingHandler()
	app.Commands = []cli.Command{
		handler.Login(),
		handler.Logout(),
	}
	return app
}

// Init Initialise a CLI app
func Init() *cli.App {
	app = cli.NewApp()
	app.Name = "orn"
	app.Usage = "orn [command]"
	app.Author = "Service API"
	app.Email = "sofyan.saputra@otoraja.co.id"
	app.Version = "0.8.1"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "config, c",
			Usage:       "Load environtment config from `FILE`",
			Destination: &Args.EnvPath,
		},
	}
	return app
}

func checkEnv() {
	if Args.EnvPath == "" {
		home, _ := os.UserHomeDir()
		err := godotenv.Load(home + "/.orn/config")
		if err != nil {
			log.Fatal("Default Configuration Not Found")
		}
	} else {
		err := godotenv.Load(Args.EnvPath)
		if err != nil {
			log.Fatal("Path Configuration Not Found")
		}
	}
}
