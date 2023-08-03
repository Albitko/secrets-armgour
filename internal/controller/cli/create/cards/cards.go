package cards

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/Albitko/secrets-armgour/internal/utils/encrypt"
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
			key, _, err := encrypt.GetCliSecrets()
			encHolder, err := encrypt.EncryptMessage([]byte(key), cardHolder)
			encNumber, err := encrypt.EncryptMessage([]byte(key), cardNumber)
			encPeriod, err := encrypt.EncryptMessage([]byte(key), cardValidityPeriod)
			encCvc, err := encrypt.EncryptMessage([]byte(key), cvcCode)
			encMeta, err := encrypt.EncryptMessage([]byte(key), meta)
			err = s.CreateCard(encHolder, encNumber, encPeriod, encCvc, encMeta)
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
