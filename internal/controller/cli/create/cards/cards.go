package cards

import (
	"github.com/spf13/cobra"
)

type sender interface {
}

func New(s sender) *cobra.Command {
	var id string
	createCmd := &cobra.Command{
		Use:   "cards",
		Short: "Create user cards secrets",
		Run: func(cmd *cobra.Command, args []string) {
			//c.sender.List
		},
	}
	createCmd.PersistentFlags().StringVarP(
		&id, "holder", "l", "", "Card holder")
	err := createCmd.MarkPersistentFlagRequired("holder")
	if err != nil {
		// TODO
		return nil
	}
	createCmd.PersistentFlags().StringVarP(
		&id, "number", "n", "", "Card number")
	err = createCmd.MarkPersistentFlagRequired("number")
	if err != nil {
		// TODO
		return nil
	}
	createCmd.PersistentFlags().StringVarP(
		&id, "period", "p", "", "Card validity period")
	err = createCmd.MarkPersistentFlagRequired("period")
	if err != nil {
		// TODO
		return nil
	}
	createCmd.PersistentFlags().StringVarP(
		&id, "cvc", "c", "", "CVC code")
	err = createCmd.MarkPersistentFlagRequired("cvc")
	if err != nil {
		// TODO
		return nil
	}
	createCmd.PersistentFlags().StringVarP(
		&id, "meta", "m", "", "Additional info about secret")
	return createCmd
}
