package binary

import (
	"fmt"

	"github.com/spf13/cobra"
)

type sender interface {
	EditBinary(index int, title, dataPath, meta string) error
}

func New(s sender) *cobra.Command {
	var title, dataPath, meta string
	var index int
	editCmd := &cobra.Command{
		Use:   "binary",
		Short: "Create user binary secret",
		Run: func(cmd *cobra.Command, args []string) {
			err := s.EditBinary(index, title, dataPath, meta)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("Binary updated successfully.")
		},
	}
	editCmd.PersistentFlags().StringVarP(
		&dataPath, "path", "p", "", "Path to binary file")
	editCmd.PersistentFlags().StringVarP(
		&title, "title", "t", "", "Title for binary secret")
	editCmd.PersistentFlags().StringVarP(
		&meta, "meta", "m", "", "Additional info about secret")
	err := editCmd.MarkPersistentFlagRequired("path")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = editCmd.MarkPersistentFlagRequired("title")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	editCmd.Flags().IntVarP(
		&index, "index", "i", 0, "index")
	return editCmd
}
