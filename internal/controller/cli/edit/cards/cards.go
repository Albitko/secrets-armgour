package cards

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/Albitko/secrets-armgour/internal/utils/encrypt"
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
			key, _, err := encrypt.GetCliSecrets()
			encHolder, err := encrypt.EncryptMessage([]byte(key), cardHolder)
			encNumber, err := encrypt.EncryptMessage([]byte(key), cardNumber)
			encPeriod, err := encrypt.EncryptMessage([]byte(key), cardValidityPeriod)
			encCvc, err := encrypt.EncryptMessage([]byte(key), cvcCode)
			encMeta, err := encrypt.EncryptMessage([]byte(key), meta)
			err = s.EditCard(index, encHolder, encNumber, encPeriod, encCvc, encMeta)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("Cards data updated successfully.")

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
