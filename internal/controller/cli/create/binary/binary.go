package binary

import (
	"fmt"

	"github.com/spf13/cobra"
)

type sender interface {
	CreateBinary(title, dataPath, meta string) error
}

// New - return command for binary creation
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
			fmt.Println("Binary data successfully written")
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
		fmt.Println(err)
		return nil
	}
	err = createCmd.MarkPersistentFlagRequired("title")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return createCmd
}
