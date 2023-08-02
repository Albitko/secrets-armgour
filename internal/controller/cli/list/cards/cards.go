package cards

import (
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	listCmd := &cobra.Command{
		Use:   "cards",
		Short: "List user saved cards secrets",
		Run: func(cmd *cobra.Command, args []string) {
			//c.sender.List
		},
	}
	return listCmd
}
