package del

import (
	"fmt"

	"github.com/spf13/cobra"
)

type sender interface {
	DeleteUserSecrets(secretType string, idx int) error
}

func New(s sender) *cobra.Command {
	var data string
	var index int
	delCmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete chosen secret",
		Run: func(cmd *cobra.Command, args []string) {
			err := s.DeleteUserSecrets(data, index)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Delete successfully")
			}
		},
	}
	delCmd.Flags().StringVarP(
		&data, "text", "t", "", "Delete text secret by index")
	delCmd.Flags().Lookup("text").NoOptDefVal = "text"

	delCmd.Flags().StringVarP(
		&data, "binary", "b", "", "Delete binary secret by index")
	delCmd.Flags().Lookup("binary").NoOptDefVal = "binary"

	delCmd.Flags().StringVarP(
		&data, "card", "c", "", "Delete card secret by index")
	delCmd.Flags().Lookup("card").NoOptDefVal = "card"

	delCmd.Flags().StringVarP(
		&data, "credentials", "r", "", "Delete credential secret by index")
	delCmd.Flags().Lookup("credentials").NoOptDefVal = "credentials"

	delCmd.Flags().IntVarP(
		&index, "index", "i", 0, "index")

	return delCmd
}
