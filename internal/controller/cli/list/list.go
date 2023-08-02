package list

import (
	"github.com/spf13/cobra"

	"github.com/Albitko/secrets-armgour/internal/controller/cli/list/all"
	"github.com/Albitko/secrets-armgour/internal/controller/cli/list/binary"
	"github.com/Albitko/secrets-armgour/internal/controller/cli/list/cards"
	"github.com/Albitko/secrets-armgour/internal/controller/cli/list/credentials"
	"github.com/Albitko/secrets-armgour/internal/controller/cli/list/text"
)

func New() *cobra.Command {
	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List user saved secrets",
		Run: func(cmd *cobra.Command, args []string) {
			//c.sender.List
		},
	}
	listCmd.AddCommand(credentials.New())
	listCmd.AddCommand(all.New())
	listCmd.AddCommand(binary.New())
	listCmd.AddCommand(cards.New())
	listCmd.AddCommand(text.New())
	return listCmd
}
