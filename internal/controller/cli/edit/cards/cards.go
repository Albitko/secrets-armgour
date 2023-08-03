package cards

import (
	"fmt"

	"github.com/spf13/cobra"
)

type sender interface {
	EditCard(index int, cardHolder, cardNumber, cardValidityPeriod, cvcCode, meta string) error
}

func New(s sender) *cobra.Command {
	var cardHolder, cardNumber, cardValidityPeriod, cvcCode, meta string
	var index int
	editCmd := &cobra.Command{
		Use:   "cards",
		Short: "Edit user cards secrets",
		Run: func(cmd *cobra.Command, args []string) {
			err := s.EditCard(index, cardHolder, cardNumber, cardValidityPeriod, cvcCode, meta)
			if err != nil {
				fmt.Println(err)
			}
		},
	}
	editCmd.PersistentFlags().StringVarP(
		&cardHolder, "holder", "l", "", "Card holder")
	err := editCmd.MarkPersistentFlagRequired("holder")
	if err != nil {
		// TODO
		return nil
	}
	editCmd.PersistentFlags().StringVarP(
		&cardNumber, "number", "n", "", "Card number")
	err = editCmd.MarkPersistentFlagRequired("number")
	if err != nil {
		// TODO
		return nil
	}
	editCmd.PersistentFlags().StringVarP(
		&cardValidityPeriod, "period", "p", "", "Card validity period")
	err = editCmd.MarkPersistentFlagRequired("period")
	if err != nil {
		// TODO
		return nil
	}
	editCmd.PersistentFlags().StringVarP(
		&cvcCode, "cvc", "c", "", "CVC code")
	err = editCmd.MarkPersistentFlagRequired("cvc")
	if err != nil {
		// TODO
		return nil
	}
	editCmd.PersistentFlags().StringVarP(
		&meta, "meta", "m", "", "Additional info about secret")
	editCmd.Flags().IntVarP(
		&index, "index", "i", 0, "index")
	return editCmd
}
