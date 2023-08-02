package edit

import (
	"github.com/spf13/cobra"

	"github.com/Albitko/secrets-armgour/internal/controller/cli/edit/binary"
	"github.com/Albitko/secrets-armgour/internal/controller/cli/edit/cards"
	"github.com/Albitko/secrets-armgour/internal/controller/cli/edit/credentials"
	"github.com/Albitko/secrets-armgour/internal/controller/cli/edit/text"
)

func New() *cobra.Command {
	editCmd := &cobra.Command{
		Use:   "edit",
		Short: "Edit secret data",
		Run: func(cmd *cobra.Command, args []string) {
			//c.sender.List
		},
	}
	editCmd.AddCommand(credentials.New())
	editCmd.AddCommand(binary.New())
	editCmd.AddCommand(cards.New())
	editCmd.AddCommand(text.New())
	return editCmd
}
