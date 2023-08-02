package binary

import (
	"github.com/spf13/cobra"
)

type sender interface {
}

func New(s sender) *cobra.Command {
	var id string
	createCmd := &cobra.Command{
		Use:   "binary",
		Short: "Create user binary secret",
		Run: func(cmd *cobra.Command, args []string) {
			//c.sender.List
		},
	}
	createCmd.PersistentFlags().StringVarP(
		&id, "path", "p", "", "Path to binary file")
	createCmd.PersistentFlags().StringVarP(
		&id, "title", "t", "", "Title for binary secret")
	createCmd.PersistentFlags().StringVarP(
		&id, "meta", "m", "", "Additional info about secret")
	err := createCmd.MarkPersistentFlagRequired("path")
	if err != nil {
		// TODO
		return nil
	}
	err = createCmd.MarkPersistentFlagRequired("title")
	if err != nil {
		// TODO
		return nil
	}
	return createCmd
}
