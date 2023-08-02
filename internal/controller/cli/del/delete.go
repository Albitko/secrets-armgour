package del

import (
	"github.com/spf13/cobra"

	"github.com/Albitko/secrets-armgour/internal/controller/cli/del/binary"
	"github.com/Albitko/secrets-armgour/internal/controller/cli/del/cards"
	"github.com/Albitko/secrets-armgour/internal/controller/cli/del/credentials"
	"github.com/Albitko/secrets-armgour/internal/controller/cli/del/text"
)

func New() *cobra.Command {
	delCmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete chosen secret",
		Run: func(cmd *cobra.Command, args []string) {
			//c.sender.List
		},
	}
	delCmd.AddCommand(credentials.New())
	delCmd.AddCommand(binary.New())
	delCmd.AddCommand(cards.New())
	delCmd.AddCommand(text.New())
	return delCmd
}
