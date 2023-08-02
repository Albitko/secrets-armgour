package text

import (
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	var id string
	editCmd := &cobra.Command{
		Use:   "text",
		Short: "Save user text secrets",
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
		&id, "title", "t", "", "Text title")
	editCmd.PersistentFlags().StringVarP(
		&id, "body", "b", "", "Text body")
	editCmd.PersistentFlags().StringVarP(
		&id, "meta", "m", "", "Additional info about secret")
	return editCmd
}
