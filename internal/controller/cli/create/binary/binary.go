package binary

import (
	"fmt"

	"github.com/spf13/cobra"
)

type sender interface {
	CreateBinary(title, dataPath, meta string) error
}

func New(s sender) *cobra.Command {
	var title, dataPath, meta string
	createCmd := &cobra.Command{
		Use:   "binary",
		Short: "Create user binary secret",
		Run: func(cmd *cobra.Command, args []string) {
			err := s.CreateBinary(title, dataPath, meta)
			if err != nil {
				fmt.Println(err)
			}
		},
	}
	createCmd.PersistentFlags().StringVarP(
		&dataPath, "path", "p", "", "Path to binary file")
	createCmd.PersistentFlags().StringVarP(
		&title, "title", "t", "", "Title for binary secret")
	createCmd.PersistentFlags().StringVarP(
		&meta, "meta", "m", "", "Additional info about secret")
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
