package credentials

import (
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	var id string
	getCmd := &cobra.Command{
		Use:   "credentials",
		Short: "Get user saved credentials. Use with ID",
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
