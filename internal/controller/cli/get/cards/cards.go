package cards

import (
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	var id string
	getCmd := &cobra.Command{
		Use:   "cards",
		Short: "Get user saved cards secrets. Use with ID",
		Run: func(cmd *cobra.Command, args []string) {
			//c.sender.List
		},
	}
	getCmd.PersistentFlags().StringVarP(
		&id, "id", "i", "", "Secret ID")
	err := getCmd.MarkPersistentFlagRequired("id")
	if err != nil {
		// TODO
		return nil
	}
	return getCmd
}
