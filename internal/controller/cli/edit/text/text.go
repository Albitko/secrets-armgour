package text

import (
	"fmt"

	"github.com/spf13/cobra"
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
			err := s.EditText(index, title, body, meta)
			if err != nil {
				fmt.Println(err)
			}
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
