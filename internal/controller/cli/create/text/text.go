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
			key, user, err := encrypt.GetCliSecrets()
			encTitle, err := encrypt.EncryptMessage([]byte(key), title)
			encBody, err := encrypt.EncryptMessage([]byte(key), body)
			encMeta, err := encrypt.EncryptMessage([]byte(key), meta)
			err = s.CreateText(encTitle, encBody, encMeta, user)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("Text data successfully written")
		},
	}

	createCmd.PersistentFlags().StringVarP(
		&title, "title", "t", "", "Text title")
	err := createCmd.MarkPersistentFlagRequired("title")
	createCmd.PersistentFlags().StringVarP(
		&body, "body", "b", "", "Text body")
	err = createCmd.MarkPersistentFlagRequired("body")
	createCmd.PersistentFlags().StringVarP(
		&meta, "meta", "m", "", "Additional info about secret")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return createCmd
}
