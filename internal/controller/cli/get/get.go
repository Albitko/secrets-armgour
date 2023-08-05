package get

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/Albitko/secrets-armgour/internal/entity"
	"github.com/Albitko/secrets-armgour/internal/utils/encrypt"
)

type sender interface {
	GetUserSecrets(secretType, user string, idx int) (interface{}, error)
}

func New(s sender) *cobra.Command {
	var data string
	var index int
	getCmd := &cobra.Command{
		Use:   "get",
		Short: "Get user saved secrets. Usage get --binary -i 23",
		Run: func(cmd *cobra.Command, args []string) {
			key, user, err := encrypt.GetCliSecrets()

			secrets, err := s.GetUserSecrets(data, user, index)
			if err != nil {
				fmt.Println(err, secrets)
			}
			switch data {
			case "credentials":
				credentials := secrets.(entity.UserCredentials)
				decMeta, err := encrypt.DecryptMessage([]byte(key), credentials.Meta)
				decService, err := encrypt.DecryptMessage([]byte(key), credentials.ServiceName)
				decLogin, err := encrypt.DecryptMessage([]byte(key), credentials.ServiceLogin)
				decPass, err := encrypt.DecryptMessage([]byte(key), credentials.ServicePassword)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println(decMeta, decService, decLogin, decPass)
			case "binary":
				bin := secrets.(entity.UserBinary)
				decMeta, err := encrypt.DecryptMessage([]byte(key), bin.Meta)
				decTitle, err := encrypt.DecryptMessage([]byte(key), bin.Title)
				decContent, err := encrypt.DecryptMessage([]byte(key), bin.B64Content)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println("Title:", decTitle, "Description:", decMeta)

				rawDecoded, err := base64.StdEncoding.DecodeString(decContent)
				if err != nil {
					fmt.Println(err)
				}
				err = os.WriteFile(decTitle+".bin", rawDecoded, 0644)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println(
					"Binary with name " + decTitle + " and meta " + decMeta + " saved as " + decTitle + ".bin")
			case "text":
				text := secrets.(entity.UserText)

				decMeta, err := encrypt.DecryptMessage([]byte(key), text.Meta)
				decTitle, err := encrypt.DecryptMessage([]byte(key), text.Title)
				decBody, err := encrypt.DecryptMessage([]byte(key), text.Body)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println(decTitle, decBody, decMeta)
			case "card":
				card := secrets.(entity.UserCard)

				holderDec, err := encrypt.DecryptMessage([]byte(key), card.CardHolder)
				numberDec, err := encrypt.DecryptMessage([]byte(key), card.CardNumber)
				periodDec, err := encrypt.DecryptMessage([]byte(key), card.CardValidityPeriod)
				cvcDec, err := encrypt.DecryptMessage([]byte(key), card.CvcCode)
				metaDec, err := encrypt.DecryptMessage([]byte(key), card.Meta)

				if err != nil {
					fmt.Println(err)
				}
				fmt.Println(holderDec, numberDec, periodDec, cvcDec, metaDec)
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
