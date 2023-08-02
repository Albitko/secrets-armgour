package credentials

import (
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	listCmd := &cobra.Command{
		Use:   "credentials",
		Short: "List user saved credentials",
		Run: func(cmd *cobra.Command, args []string) {
			//c.sender.List
		},
	}
	return listCmd
}
