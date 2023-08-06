package cards

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/Albitko/secrets-armgour/internal/utils/encrypt"
)

type sender interface {
	EditCard(index int, cardHolder, cardNumber, cardValidityPeriod, cvcCode, meta string) error
}

// New - return command for cards edit
func New(s sender) *cobra.Command {
	var cardHolder, cardNumber, cardValidityPeriod, cvcCode, meta string
	var index int
	editCmd := &cobra.Command{
		Use:   "cards",
		Short: "Edit user cards secrets",
		Run: func(cmd *cobra.Command, args []string) {
			var key, encHolder, encNumber, encPeriod, encCvc, encMeta string
			var err error
			key, _, err = encrypt.GetCliSecrets()
			if err != nil {
				fmt.Println(err)
				return
			}
			encHolder, err = encrypt.EncryptMessage([]byte(key), cardHolder)
			if err != nil {
				fmt.Println(err)
				return
			}
			encNumber, err = encrypt.EncryptMessage([]byte(key), cardNumber)
			if err != nil {
				fmt.Println(err)
				return
			}
			encPeriod, err = encrypt.EncryptMessage([]byte(key), cardValidityPeriod)
			if err != nil {
				fmt.Println(err)
				return
			}
			encCvc, err = encrypt.EncryptMessage([]byte(key), cvcCode)
			if err != nil {
				fmt.Println(err)
				return
			}
			encMeta, err = encrypt.EncryptMessage([]byte(key), meta)
			if err != nil {
				fmt.Println(err)
				return
			}
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
		fmt.Println(err)
		return nil
	}
	editCmd.PersistentFlags().StringVarP(
		&cardNumber, "number", "n", "", "Card number")
	err = editCmd.MarkPersistentFlagRequired("number")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	editCmd.PersistentFlags().StringVarP(
		&cardValidityPeriod, "period", "p", "", "Card validity period")
	err = editCmd.MarkPersistentFlagRequired("period")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	editCmd.PersistentFlags().StringVarP(
		&cvcCode, "cvc", "c", "", "CVC code")
	err = editCmd.MarkPersistentFlagRequired("cvc")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	editCmd.PersistentFlags().StringVarP(
		&meta, "meta", "m", "", "Additional info about secret")
	editCmd.Flags().IntVarP(
		&index, "index", "i", 0, "index")
	return editCmd
}
