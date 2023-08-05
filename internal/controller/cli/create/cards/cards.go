package cards

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/Albitko/secrets-armgour/internal/utils/encrypt"
)

type sender interface {
	CreateCard(cardHolder, cardNumber, cardValidityPeriod, cvcCode, meta, user string) error
}

func New(s sender) *cobra.Command {
	var cardHolder, cardNumber, cardValidityPeriod, cvcCode, meta string
	createCmd := &cobra.Command{
		Use:   "cards",
		Short: "Create user cards secrets",
		Run: func(cmd *cobra.Command, args []string) {
			var encHolder, encNumber, encPeriod string
			var err error
			key, user, err := encrypt.GetCliSecrets()
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
			encCvc, err := encrypt.EncryptMessage([]byte(key), cvcCode)
			if err != nil {
				fmt.Println(err)
				return
			}
			encMeta, err := encrypt.EncryptMessage([]byte(key), meta)
			if err != nil {
				fmt.Println(err)
				return
			}
			err = s.CreateCard(encHolder, encNumber, encPeriod, encCvc, encMeta, user)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("Card data successfully written")
		},
	}
	createCmd.PersistentFlags().StringVarP(
		&cardHolder, "holder", "l", "", "Card holder")
	err := createCmd.MarkPersistentFlagRequired("holder")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	createCmd.PersistentFlags().StringVarP(
		&cardNumber, "number", "n", "", "Card number")
	err = createCmd.MarkPersistentFlagRequired("number")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	createCmd.PersistentFlags().StringVarP(
		&cardValidityPeriod, "period", "p", "", "Card validity period")
	err = createCmd.MarkPersistentFlagRequired("period")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	createCmd.PersistentFlags().StringVarP(
		&cvcCode, "cvc", "c", "", "CVC code")
	err = createCmd.MarkPersistentFlagRequired("cvc")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	createCmd.PersistentFlags().StringVarP(
		&meta, "meta", "m", "", "Additional info about secret")
	return createCmd
}
