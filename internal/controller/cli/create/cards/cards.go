package cards

import (
	"fmt"

	"github.com/spf13/cobra"
)

type sender interface {
	CreateCard(cardHolder, cardNumber, cardValidityPeriod, cvcCode, meta string) error
}

func New(s sender) *cobra.Command {
	var cardHolder, cardNumber, cardValidityPeriod, cvcCode, meta string
	createCmd := &cobra.Command{
		Use:   "cards",
		Short: "Create user cards secrets",
		Run: func(cmd *cobra.Command, args []string) {
			err := s.CreateCard(cardHolder, cardNumber, cardValidityPeriod, cvcCode, meta)
			if err != nil {
				fmt.Println(err)
			}
		},
	}
	createCmd.PersistentFlags().StringVarP(
		&cardHolder, "holder", "l", "", "Card holder")
	err := createCmd.MarkPersistentFlagRequired("holder")
	if err != nil {
		// TODO
		return nil
	}
	createCmd.PersistentFlags().StringVarP(
		&cardNumber, "number", "n", "", "Card number")
	err = createCmd.MarkPersistentFlagRequired("number")
	if err != nil {
		// TODO
		return nil
	}
	createCmd.PersistentFlags().StringVarP(
		&cardValidityPeriod, "period", "p", "", "Card validity period")
	err = createCmd.MarkPersistentFlagRequired("period")
	if err != nil {
		// TODO
		return nil
	}
	createCmd.PersistentFlags().StringVarP(
		&cvcCode, "cvc", "c", "", "CVC code")
	err = createCmd.MarkPersistentFlagRequired("cvc")
	if err != nil {
		// TODO
		return nil
	}
	createCmd.PersistentFlags().StringVarP(
		&meta, "meta", "m", "", "Additional info about secret")
	return createCmd
}
