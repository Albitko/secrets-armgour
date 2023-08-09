package edit

import (
	"github.com/spf13/cobra"

	"github.com/Albitko/secrets-armgour/internal/controller/cli/edit/binary"
	"github.com/Albitko/secrets-armgour/internal/controller/cli/edit/cards"
	"github.com/Albitko/secrets-armgour/internal/controller/cli/edit/credentials"
	"github.com/Albitko/secrets-armgour/internal/controller/cli/edit/text"
)

type sender interface {
	EditCredentials(index int, serviceName, serviceLogin, servicePassword, meta string) error
	EditBinary(index int, title, dataPath, meta string) error
	EditCard(index int, cardHolder, cardNumber, cardValidityPeriod, cvcCode, meta string) error
	EditText(index int, title, body, meta string) error
}

// New - return command for data edit
func New(s sender) *cobra.Command {
	editCmd := &cobra.Command{
		Use:   "edit",
		Short: "Edit secret data",
	}
	editCmd.AddCommand(credentials.New(s))
	editCmd.AddCommand(binary.New(s))
	editCmd.AddCommand(cards.New(s))
	editCmd.AddCommand(text.New(s))
	return editCmd
}
