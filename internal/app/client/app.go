package client

import (
	"fmt"
	"os"

	"github.com/Albitko/secrets-armgour/internal/adapter/httpapi"
	"github.com/Albitko/secrets-armgour/internal/controller/cli"
	"github.com/Albitko/secrets-armgour/internal/secrets/sender"
)

func Run() {
	serviceAPI := httpapi.New()
	secretsSender := sender.New(serviceAPI)
	clientCli := cli.New(secretsSender)

	rootCmd := clientCli.Cmd()
	rootCmd.AddCommand(clientCli.Login())
	rootCmd.AddCommand(clientCli.Logout())
	rootCmd.AddCommand(clientCli.Register())
	rootCmd.AddCommand(clientCli.List())
	rootCmd.AddCommand(clientCli.Get())
	rootCmd.AddCommand(clientCli.Create())
	rootCmd.AddCommand(clientCli.Edit())
	rootCmd.AddCommand(clientCli.Delete())
	rootCmd.AddCommand(clientCli.GeneratePassword())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
