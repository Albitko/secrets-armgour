package list

import (
	"github.com/spf13/cobra"

	"github.com/Albitko/secrets-armgour/internal/controller/cli/list/binary"
	"github.com/Albitko/secrets-armgour/internal/controller/cli/list/cards"
	"github.com/Albitko/secrets-armgour/internal/controller/cli/list/credentials"
	"github.com/Albitko/secrets-armgour/internal/controller/cli/list/text"
)

type sender interface {
	GetAllSecrets()
}

func New(s sender) *cobra.Command {
	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List user saved secrets",
		Run: func(cmd *cobra.Command, args []string) {
			//c.sender.List
		},
	}
	listCmd.AddCommand(credentials.New(s))
	listCmd.AddCommand(binary.New(s))
	listCmd.AddCommand(cards.New(s))
	listCmd.AddCommand(text.New(s))
	return listCmd
}
