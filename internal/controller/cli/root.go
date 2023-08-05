package cli

import (
	"github.com/spf13/cobra"

	"github.com/Albitko/secrets-armgour/internal/controller/cli/create"
	"github.com/Albitko/secrets-armgour/internal/controller/cli/del"
	"github.com/Albitko/secrets-armgour/internal/controller/cli/edit"
	"github.com/Albitko/secrets-armgour/internal/controller/cli/get"
	"github.com/Albitko/secrets-armgour/internal/controller/cli/list"
	"github.com/Albitko/secrets-armgour/internal/controller/cli/login"
	"github.com/Albitko/secrets-armgour/internal/controller/cli/logout"
	"github.com/Albitko/secrets-armgour/internal/controller/cli/register"
)

type sender interface {
	CreateCredentials(serviceName, serviceLogin, servicePassword, meta, user string) error
	CreateText(title, body, meta, user string) error
	CreateCard(cardHolder, cardNumber, cardValidityPeriod, cvcCode, meta, user string) error
	CreateBinary(title, dataPath, meta string) error

	ListUserSecrets(data, user string) (interface{}, error)
	GetUserSecrets(secretType, user string, idx int) (interface{}, error)
	DeleteUserSecrets(secretType string, idx int) error
	EditCredentials(index int, serviceName, serviceLogin, servicePassword, meta string) error
	EditBinary(index int, title, dataPath, meta string) error
	EditCard(index int, cardHolder, cardNumber, cardValidityPeriod, cvcCode, meta string) error
	EditText(index int, title, body, meta string) error

	RegisterUser(login, password string) error
	LoginUser(login, password string) error
}

type cliCommands struct {
	s   sender
	Cmd *cobra.Command
}

func New(s sender) *cliCommands {
	rootCmd := &cobra.Command{
		Use:   "armgour-cli",
		Short: "Client for storing your secrets in armGOur service",
	}

	rootCmd.AddCommand(login.New(s))
	rootCmd.AddCommand(logout.New())
	rootCmd.AddCommand(register.New(s))
	rootCmd.AddCommand(list.New(s))
	rootCmd.AddCommand(get.New(s))
	rootCmd.AddCommand(create.New(s))
	rootCmd.AddCommand(edit.New(s))
	rootCmd.AddCommand(del.New(s))
	return &cliCommands{
		s:   s,
		Cmd: rootCmd,
	}
}
