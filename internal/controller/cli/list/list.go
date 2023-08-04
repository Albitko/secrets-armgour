package list

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/Albitko/secrets-armgour/internal/entity"
	"github.com/Albitko/secrets-armgour/internal/utils/encrypt"
)

type sender interface {
	ListUserSecrets(data, user string) (interface{}, error)
}

func New(s sender) *cobra.Command {
	var data string
	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List user saved secrets",
		Run: func(cmd *cobra.Command, args []string) {
			key, user, err := encrypt.GetCliSecrets()
			if err != nil {
				fmt.Println(err)
				return
			}

			res, err := s.ListUserSecrets(data, user)
			if err != nil {
				fmt.Println(err)
			}
			switch data {
			case "credentials":
				credentials := res.([]entity.CutCredentials)
				for _, c := range credentials {
					decMeta, err := encrypt.DecryptMessage([]byte(key), c.Meta)
					decService, err := encrypt.DecryptMessage([]byte(key), c.ServiceName)
					if err != nil {
						fmt.Println(err)
					}
					fmt.Println(c.Id, decService, decMeta)
				}
			case "binary":
				bin := res.([]entity.CutBinary)
				for _, b := range bin {
					decMeta, err := encrypt.DecryptMessage([]byte(key), b.Meta)
					decTitle, err := encrypt.DecryptMessage([]byte(key), b.Title)
					if err != nil {
						fmt.Println(err)
					}
					fmt.Println(b.Id, decTitle, decMeta)
				}
			case "text":
				texts := res.([]entity.CutText)
				for _, t := range texts {
					decMeta, err := encrypt.DecryptMessage([]byte(key), t.Meta)
					decTitle, err := encrypt.DecryptMessage([]byte(key), t.Title)
					if err != nil {
						fmt.Println(err)
					}
					fmt.Println(t.Id, decTitle, decMeta)
				}
			case "card":
				cards := res.([]entity.CutCard)
				for _, c := range cards {
					numberDec, err := encrypt.DecryptMessage([]byte(key), c.CardNumber)
					metaDec, err := encrypt.DecryptMessage([]byte(key), c.Meta)

					if err != nil {
						fmt.Println(err)
					}
					fmt.Println(c.Id, numberDec, metaDec)
				}
			}
		},
	}
	listCmd.Flags().StringVarP(
		&data, "text", "t", "text", "List text secrets")
	listCmd.Flags().Lookup("text").NoOptDefVal = "text"
	listCmd.Flags().StringVarP(
		&data, "binary", "b", "binary", "List binary secrets")
	listCmd.Flags().Lookup("binary").NoOptDefVal = "binary"
	listCmd.Flags().StringVarP(
		&data, "card", "c", "card", "List cards secrets")
	listCmd.Flags().Lookup("card").NoOptDefVal = "card"
	listCmd.Flags().StringVarP(
		&data, "credentials", "r", "credentials", "List credentials secrets")
	listCmd.Flags().Lookup("credentials").NoOptDefVal = "credentials"
	return listCmd
}
