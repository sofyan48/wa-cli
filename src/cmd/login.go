package cmd

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func (handler *CLIMapping) Login() cli.Command {
	command := cli.Command{}
	command.Name = "login"
	command.Usage = "logout [option]"
	command.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "path, p",
			Usage:       "QR path",
			Destination: &Args.QrPath,
		},
		// cli.BoolFlag{
		// 	Name:        "decryption",
		// 	Usage:       "--decryption",
		// 	Destination: &Args.Decryption,
		// },
	}
	command.Action = func(c *cli.Context) error {
		checkEnv()
		wac, err := handler.chat.Conn()
		if err != nil {
			fmt.Println("Connection Failed: ", err)
			os.Exit(1)
		}
		err = handler.chat.Login(wac)
		if err != nil {
			fmt.Println("Login Failed: ", err)
			os.Exit(1)
		}
		fmt.Println("Login Successfully")
		return nil
	}

	return command
}
