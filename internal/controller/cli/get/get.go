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

// New - return command for secrets getting
func New(s sender) *cobra.Command {
	var data string
	var index int
	getCmd := &cobra.Command{
		Use:   "get",
		Short: "Get user saved secrets. Usage get --binary -i 23",
		Run: func(cmd *cobra.Command, args []string) {
			var key, user string
			var err error
			var secrets interface{}
			key, user, err = encrypt.GetCliSecrets()
			if err != nil {
				fmt.Println(err)
			}
			secrets, err = s.GetUserSecrets(data, user, index)
			if err != nil {
				fmt.Println(err, secrets)
			}
			switch data {
			case entity.Credentials:
				credentials := secrets.(entity.UserCredentials)
				decMeta, err := encrypt.DecryptMessage([]byte(key), credentials.Meta)
				if err != nil {
					fmt.Println("Error reading file:", err)
					return
				}
				decService, err := encrypt.DecryptMessage([]byte(key), credentials.ServiceName)
				if err != nil {
					fmt.Println("Error reading file:", err)
					return
				}
				decLogin, err := encrypt.DecryptMessage([]byte(key), credentials.ServiceLogin)
				if err != nil {
					fmt.Println("Error reading file:", err)
					return
				}
				decPass, err := encrypt.DecryptMessage([]byte(key), credentials.ServicePassword)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println(
					"Service name:", decService, "Login:", decLogin, "Password:", decPass, "Description:", decMeta)
			case entity.Binary:
				bin := secrets.(entity.UserBinary)
				decMeta, err := encrypt.DecryptMessage([]byte(key), bin.Meta)
				if err != nil {
					fmt.Println("Error reading file:", err)
					return
				}
				decTitle, err := encrypt.DecryptMessage([]byte(key), bin.Title)
				if err != nil {
					fmt.Println("Error reading file:", err)
					return
				}
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
			case entity.Text:
				text := secrets.(entity.UserText)

				decMeta, err := encrypt.DecryptMessage([]byte(key), text.Meta)
				if err != nil {
					fmt.Println("Error reading file:", err)
					return
				}
				decTitle, err := encrypt.DecryptMessage([]byte(key), text.Title)
				if err != nil {
					fmt.Println("Error reading file:", err)
					return
				}
				decBody, err := encrypt.DecryptMessage([]byte(key), text.Body)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println("Note title", decTitle, "Note text", decBody, "Description:", decMeta)
			case entity.Card:
				card := secrets.(entity.UserCard)

				holderDec, err := encrypt.DecryptMessage([]byte(key), card.CardHolder)
				if err != nil {
					fmt.Println(err)
					return
				}
				numberDec, err := encrypt.DecryptMessage([]byte(key), card.CardNumber)
				if err != nil {
					fmt.Println(err)
					return
				}
				periodDec, err := encrypt.DecryptMessage([]byte(key), card.CardValidityPeriod)
				if err != nil {
					fmt.Println(err)
					return
				}
				cvcDec, err := encrypt.DecryptMessage([]byte(key), card.CvcCode)
				if err != nil {
					fmt.Println(err)
					return
				}
				metaDec, err := encrypt.DecryptMessage([]byte(key), card.Meta)
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println("Card holder", holderDec,
					"Card number", numberDec, "Validity period", periodDec, "CVC", cvcDec, "Description:", metaDec)
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
