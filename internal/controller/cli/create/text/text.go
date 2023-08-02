package text

import (
	"fmt"

	"github.com/spf13/cobra"
)

type sender interface {
	CreateText(title, body, meta string) error
}

func New(s sender) *cobra.Command {
	var title, body, meta string
	createCmd := &cobra.Command{
		Use:   "text",
		Short: "Save user text secrets",
		Run: func(cmd *cobra.Command, args []string) {
			err := s.CreateText(title, body, meta)
			if err != nil {
				fmt.Println(err)
			}
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
		// TODO
		return nil
	}
	return createCmd
}
