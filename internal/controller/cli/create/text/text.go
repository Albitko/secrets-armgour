package text

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/Albitko/secrets-armgour/internal/utils/encrypt"
)

type sender interface {
	CreateText(title, body, meta, user string) error
}

func New(s sender) *cobra.Command {
	var title, body, meta string
	createCmd := &cobra.Command{
		Use:   "text",
		Short: "Save user text secrets",
		Run: func(cmd *cobra.Command, args []string) {
			var key, user, encTitle, encBody, encMeta string
			var err error
			key, user, err = encrypt.GetCliSecrets()
			if err != nil {
				fmt.Println(err)
				return
			}
			encTitle, err = encrypt.EncryptMessage([]byte(key), title)
			if err != nil {
				fmt.Println(err)
				return
			}
			encBody, err = encrypt.EncryptMessage([]byte(key), body)
			if err != nil {
				fmt.Println(err)
				return
			}
			encMeta, err = encrypt.EncryptMessage([]byte(key), meta)
			if err != nil {
				fmt.Println(err)
				return
			}
			err = s.CreateText(encTitle, encBody, encMeta, user)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("Text data successfully written")
		},
	}

	createCmd.PersistentFlags().StringVarP(
		&title, "title", "t", "", "Text title")
	err := createCmd.MarkPersistentFlagRequired("title")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	createCmd.PersistentFlags().StringVarP(
		&body, "body", "b", "", "Text body")
	err = createCmd.MarkPersistentFlagRequired("body")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	createCmd.PersistentFlags().StringVarP(
		&meta, "meta", "m", "", "Additional info about secret")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return createCmd
}
