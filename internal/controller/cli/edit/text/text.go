package text

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/Albitko/secrets-armgour/internal/utils/encrypt"
)

type sender interface {
	EditText(index int, title, body, meta string) error
}

func New(s sender) *cobra.Command {
	var title, body, meta string
	var index int
	editCmd := &cobra.Command{
		Use:   "text",
		Short: "Edit user text secrets",
		Run: func(cmd *cobra.Command, args []string) {
			key, _, err := encrypt.GetCliSecrets()
			encTitle, err := encrypt.EncryptMessage([]byte(key), title)
			encBody, err := encrypt.EncryptMessage([]byte(key), body)
			encMeta, err := encrypt.EncryptMessage([]byte(key), meta)
			err = s.EditText(index, encTitle, encBody, encMeta)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("Text data updated successfully.")
		},
	}
	editCmd.PersistentFlags().StringVarP(
		&title, "title", "t", "", "Text title")
	err := editCmd.MarkPersistentFlagRequired("title")
	editCmd.PersistentFlags().StringVarP(
		&body, "body", "b", "", "Text body")
	err = editCmd.MarkPersistentFlagRequired("body")
	editCmd.PersistentFlags().StringVarP(
		&meta, "meta", "m", "", "Additional info about secret")
	if err != nil {
		// TODO
		return nil
	}
	editCmd.Flags().IntVarP(
		&index, "index", "i", 0, "index")
	return editCmd
}
