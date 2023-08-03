package get

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/Albitko/secrets-armgour/internal/entity"
)

type sender interface {
	GetUserSecrets(secretType string, idx int) (interface{}, error)
}

func New(s sender) *cobra.Command {
	var data string
	var index int
	getCmd := &cobra.Command{
		Use:   "get",
		Short: "Get user saved secrets. Usage get --binary -i 23",
		Run: func(cmd *cobra.Command, args []string) {
			secrets, err := s.GetUserSecrets(data, index)
			if err != nil {
				fmt.Println(err, secrets)
			}
			switch data {
			case "credentials":
				credentials := secrets.(entity.UserCredentials)
				fmt.Println(credentials)
			case "binary":
				bin := secrets.(entity.UserBinary)
				fmt.Println(bin)
				fmt.Println("Binary saved in .. with name ..")
			case "text":
				text := secrets.(entity.UserText)
				fmt.Println(text)
			case "card":
				card := secrets.(entity.UserCard)
				fmt.Println(card)
			}
		},
	}
	getCmd.Flags().StringVarP(
		&data, "text", "t", "", "Get text secret by index")
	getCmd.Flags().Lookup("text").NoOptDefVal = "text"

	getCmd.Flags().StringVarP(
		&data, "binary", "b", "", "Get binary secret by index")
	getCmd.Flags().Lookup("binary").NoOptDefVal = "binary"

	getCmd.Flags().StringVarP(
		&data, "card", "c", "", "Get card secret by index")
	getCmd.Flags().Lookup("card").NoOptDefVal = "card"

	getCmd.Flags().StringVarP(
		&data, "credentials", "r", "", "Get credential secret by index")
	getCmd.Flags().Lookup("credentials").NoOptDefVal = "credentials"

	getCmd.Flags().IntVarP(
		&index, "index", "i", 0, "index")

	return getCmd
}
