package text

import (
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	var id string
	delCmd := &cobra.Command{
		Use:   "text",
		Short: "Delete user saved text secrets. Use with ID",
		Run: func(cmd *cobra.Command, args []string) {
			//c.sender.List
		},
	}
	delCmd.PersistentFlags().StringVarP(
		&id, "id", "i", "", "Secret ID")
	err := delCmd.MarkPersistentFlagRequired("id")
	if err != nil {
		// TODO
		return nil
	}
	return delCmd
}
