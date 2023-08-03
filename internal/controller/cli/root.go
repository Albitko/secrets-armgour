package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/Albitko/secrets-armgour/internal/controller/cli/create"
	"github.com/Albitko/secrets-armgour/internal/controller/cli/del"
	"github.com/Albitko/secrets-armgour/internal/controller/cli/edit"
	"github.com/Albitko/secrets-armgour/internal/controller/cli/gen"
	"github.com/Albitko/secrets-armgour/internal/controller/cli/get"
	"github.com/Albitko/secrets-armgour/internal/controller/cli/list"
	"github.com/Albitko/secrets-armgour/internal/controller/cli/login"
	"github.com/Albitko/secrets-armgour/internal/controller/cli/logout"
	"github.com/Albitko/secrets-armgour/internal/controller/cli/register"
)

type sender interface {
	CreateCredentials(serviceName, serviceLogin, servicePassword, meta string) error
	CreateText(title, body, meta string) error
	CreateCard(cardHolder, cardNumber, cardValidityPeriod, cvcCode, meta string) error
	CreateBinary(title, dataPath, meta string) error

	ListUserSecrets(data string) (interface{}, error)
	GetUserSecrets(secretType string, idx int) (interface{}, error)
	DeleteUserSecrets(secretType string, idx int) error
	EditCredentials(index int, serviceName, serviceLogin, servicePassword, meta string) error
	EditBinary(index int, title, dataPath, meta string) error
	EditCard(index int, cardHolder, cardNumber, cardValidityPeriod, cvcCode, meta string) error
	EditText(index int, title, body, meta string) error
}

type cliCommands struct {
	isAuth bool
	s      sender
	Cmd    *cobra.Command
}

func New(s sender) *cliCommands {
	isAuth := false
	rootCmd := &cobra.Command{
		Use:   "armgour-cli",
		Short: "Client for storing your secrets in armGOur service",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			// TODO mock token check
			// check if no token
			content, err := os.ReadFile(".token")
			if err != nil {
				if os.IsNotExist(err) {
					fmt.Println("AUTH please")
				} else {
					fmt.Println("Error reading file:", err)
				}
				return
			}
			fmt.Println("isAuth", isAuth, content)
			// or check if token expired
			isAuth = true
		},
	}
	rootCmd.AddCommand(login.New())
	rootCmd.AddCommand(logout.New())
	rootCmd.AddCommand(register.New())
	rootCmd.AddCommand(list.New(s))
	rootCmd.AddCommand(get.New(s))
	rootCmd.AddCommand(create.New(s))
	rootCmd.AddCommand(edit.New(s))
	rootCmd.AddCommand(del.New(s))
	rootCmd.AddCommand(gen.New())
	return &cliCommands{
		s:      s,
		Cmd:    rootCmd,
		isAuth: isAuth,
	}
}
