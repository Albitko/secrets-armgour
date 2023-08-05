package create

import (
	"github.com/spf13/cobra"

	"github.com/Albitko/secrets-armgour/internal/controller/cli/create/binary"
	"github.com/Albitko/secrets-armgour/internal/controller/cli/create/cards"
	"github.com/Albitko/secrets-armgour/internal/controller/cli/create/credentials"
	"github.com/Albitko/secrets-armgour/internal/controller/cli/create/text"
)

type sender interface {
	CreateCredentials(serviceName, serviceLogin, servicePassword, meta, user string) error
	CreateText(title, body, meta, user string) error
	CreateCard(cardHolder, cardNumber, cardValidityPeriod, cvcCode, meta, user string) error
	CreateBinary(title, dataPath, meta string) error
}

func New(s sender) *cobra.Command {
	createCmd := &cobra.Command{
		Use:   "create",
		Short: "Create secret",
	}
	createCmd.AddCommand(credentials.New(s))
	createCmd.AddCommand(binary.New(s))
	createCmd.AddCommand(cards.New(s))
	createCmd.AddCommand(text.New(s))
	return createCmd
}
