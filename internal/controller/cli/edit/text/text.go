package text

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/Albitko/secrets-armgour/internal/utils/encrypt"
)

type sender interface {
	EditText(index int, title, body, meta string) error
}

// New - return command for all secret data deletion
func New(s sender) *cobra.Command {
	var title, body, meta string
	var index int
	editCmd := &cobra.Command{
		Use:   "text",
		Short: "Edit user text secrets",
		Run: func(cmd *cobra.Command, args []string) {
			var key, encTitle, encBody, encMeta string
			var err error
			key, _, err = encrypt.GetCliSecrets()
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
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}
	editCmd.PersistentFlags().StringVarP(
		&body, "body", "b", "", "Text body")
	err = editCmd.MarkPersistentFlagRequired("body")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}
	editCmd.PersistentFlags().StringVarP(
		&meta, "meta", "m", "", "Additional info about secret")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	editCmd.Flags().IntVarP(
		&index, "index", "i", 0, "index")
	return editCmd
}
