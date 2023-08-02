package binary

import (
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	var id string
	editCmd := &cobra.Command{
		Use:   "binary",
		Short: "Edit user binary secret",
		Run: func(cmd *cobra.Command, args []string) {
			//c.sender.List
		},
	}
	editCmd.PersistentFlags().StringVarP(
		&id, "id", "i", "", "Secret ID")
	err := editCmd.MarkPersistentFlagRequired("id")
	if err != nil {
		// TODO
		return nil
	}
	editCmd.PersistentFlags().StringVarP(
		&id, "path", "p", "", "Path to binary file")
	editCmd.PersistentFlags().StringVarP(
		&id, "title", "t", "", "Title for binary secret")
	editCmd.PersistentFlags().StringVarP(
		&id, "meta", "m", "", "Additional info about secret")
	return editCmd
}
