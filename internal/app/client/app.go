package client

import (
	"flag"
	"fmt"
	"os"

	"github.com/Albitko/secrets-armgour/internal/adapter/httpapi"
	"github.com/Albitko/secrets-armgour/internal/controller/cli"
	"github.com/Albitko/secrets-armgour/internal/secrets/sender"
)

func Run() {
	serviceURL := flag.String("address", "http://localhost:8080", "Address of armGour service")
	flag.Parse()
	serviceAPI := httpapi.New(*serviceURL)
	secretsSender := sender.New(serviceAPI)
	c := cli.New(secretsSender)

	if err := c.Cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
