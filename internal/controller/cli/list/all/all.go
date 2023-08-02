package all

import (
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	listCmd := &cobra.Command{
		Use:   "all",
		Short: "List all user saved secrets",
		Run: func(cmd *cobra.Command, args []string) {
			//c.sender.List
		},
	}
	return listCmd
}
