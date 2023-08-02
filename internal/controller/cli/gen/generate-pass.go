package gen

import (
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	genCmd := &cobra.Command{
		Use:   "gen",
		Short: "Generate random password",
		Run: func(cmd *cobra.Command, args []string) {
			//c.sender.List
		},
	}
	return genCmd
}
