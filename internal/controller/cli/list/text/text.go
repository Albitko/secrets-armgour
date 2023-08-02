package text

import (
	"github.com/spf13/cobra"
)

type sender interface {
}

func New(s sender) *cobra.Command {
	listCmd := &cobra.Command{
		Use:   "text",
		Short: "List user saved text secrets",
		Run: func(cmd *cobra.Command, args []string) {
			//c.sender.List
		},
	}
	return listCmd
}
