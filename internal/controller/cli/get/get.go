package get

import (
	"github.com/spf13/cobra"

	"github.com/Albitko/secrets-armgour/internal/controller/cli/get/binary"
	"github.com/Albitko/secrets-armgour/internal/controller/cli/get/cards"
	"github.com/Albitko/secrets-armgour/internal/controller/cli/get/credentials"
	"github.com/Albitko/secrets-armgour/internal/controller/cli/get/text"
)

func New() *cobra.Command {
	getCmd := &cobra.Command{
		Use:   "get",
		Short: "Get user saved secrets",
		Run: func(cmd *cobra.Command, args []string) {
			//c.sender.List
		},
	}
	getCmd.AddCommand(credentials.New())
	getCmd.AddCommand(binary.New())
	getCmd.AddCommand(cards.New())
	getCmd.AddCommand(text.New())
	return getCmd
}
