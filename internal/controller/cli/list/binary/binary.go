package binary

import (
	"github.com/spf13/cobra"
)

type sender interface {
}

func New(s sender) *cobra.Command {
	listCmd := &cobra.Command{
		Use:   "binary",
		Short: "List user binary secrets",
		Run: func(cmd *cobra.Command, args []string) {
			//c.sender.List
		},
	}
	return listCmd
}
