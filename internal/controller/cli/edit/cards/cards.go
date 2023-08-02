package cards

import (
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	var id string
	editCmd := &cobra.Command{
		Use:   "cards",
		Short: "Create user cards secrets",
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
		&id, "holder", "l", "", "Card holder")
	editCmd.PersistentFlags().StringVarP(
		&id, "number", "n", "", "Card number")
	editCmd.PersistentFlags().StringVarP(
		&id, "period", "p", "", "Card validity period")
	editCmd.PersistentFlags().StringVarP(
		&id, "cvc", "c", "", "CVC code")
	editCmd.PersistentFlags().StringVarP(
		&id, "meta", "m", "", "Additional info about secret")
	return editCmd
}
