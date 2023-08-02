package client

import (
	"fmt"
	"os"

	"github.com/Albitko/secrets-armgour/internal/adapter/httpapi"
	"github.com/Albitko/secrets-armgour/internal/controller/cli"
	"github.com/Albitko/secrets-armgour/internal/secrets/sender"
)

func Run() {
	serviceAPI := httpapi.New("http://localhost:8080")
	secretsSender := sender.New(serviceAPI)
	c := cli.New(secretsSender)

	if err := c.Cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
