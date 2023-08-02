package text

import (
	"github.com/spf13/cobra"
)

type sender interface {
}

func New(s sender) *cobra.Command {
	var id string
	createCmd := &cobra.Command{
		Use:   "text",
		Short: "Save user text secrets",
		Run: func(cmd *cobra.Command, args []string) {
			//c.sender.List
		},
	}
	createCmd.PersistentFlags().StringVarP(
		&id, "title", "t", "", "Text title")
	err := createCmd.MarkPersistentFlagRequired("title")
	createCmd.PersistentFlags().StringVarP(
		&id, "body", "b", "", "Text body")
	err = createCmd.MarkPersistentFlagRequired("body")
	createCmd.PersistentFlags().StringVarP(
		&id, "meta", "m", "", "Additional info about secret")
	if err != nil {
		// TODO
		return nil
	}
	return createCmd
}
