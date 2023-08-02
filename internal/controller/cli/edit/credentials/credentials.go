package credentials

import (
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	var id string
	editCmd := &cobra.Command{
		Use:   "credentials",
		Short: "Save user credentials",
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
		&id, "service", "s", "", "Service name")
	editCmd.PersistentFlags().StringVarP(
		&id, "login", "l", "", "Service login")
	editCmd.PersistentFlags().StringVarP(
		&id, "password", "p", "", "Service password")
	editCmd.PersistentFlags().StringVarP(
		&id, "meta", "m", "", "Additional info about secret")
	return editCmd
}
