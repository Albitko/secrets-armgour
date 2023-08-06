package client

import (
	"flag"
	"fmt"
	"os"

	"github.com/Albitko/secrets-armgour/internal/adapter/httpapi"
	"github.com/Albitko/secrets-armgour/internal/controller/cli"
	"github.com/Albitko/secrets-armgour/internal/secrets/sender"
)

// Run - client application
func Run() {
	serviceURL := flag.String("address", "http://localhost:8080", "Address of armGour service")
	version := flag.String("version", "0.1.0", "Version of armGour service")
	buildAt := flag.String("date", "05.08.2023", "Date of build")
	flag.Parse()
	fmt.Println("Secrets armGour version", *version)
	fmt.Println("Build at", *buildAt)
	fmt.Println("===================================")

	serviceAPI := httpapi.New(*serviceURL)
	secretsSender := sender.New(serviceAPI)
	c := cli.New(secretsSender)

	if err := c.Cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
