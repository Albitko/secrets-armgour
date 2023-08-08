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

// New - return command for list secrets
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
			case entity.Credentials:
				credentials := res.([]entity.CutCredentials)
				for _, c := range credentials {
					decMeta, err := encrypt.DecryptMessage([]byte(key), c.Meta)
					if err != nil {
						fmt.Println(err)
						return
					}
					decService, err := encrypt.DecryptMessage([]byte(key), c.ServiceName)
					if err != nil {
						fmt.Println(err)
						return
					}
					fmt.Println("ID:", c.Id, "Service name:", decService, "Description:", decMeta)
				}
			case entity.Binary:
				bin := res.([]entity.CutBinary)
				for _, b := range bin {
					decMeta, err := encrypt.DecryptMessage([]byte(key), b.Meta)
					if err != nil {
						fmt.Println(err)
						return
					}
					decTitle, err := encrypt.DecryptMessage([]byte(key), b.Title)
					if err != nil {
						fmt.Println(err)
						return
					}
					fmt.Println("ID:", b.Id, "Binary name:", decTitle, "Description:", decMeta)
				}
			case entity.Text:
				texts := res.([]entity.CutText)
				for _, t := range texts {
					decMeta, err := encrypt.DecryptMessage([]byte(key), t.Meta)
					if err != nil {
						fmt.Println(err)
						return
					}
					decTitle, err := encrypt.DecryptMessage([]byte(key), t.Title)
					if err != nil {
						fmt.Println(err)
						return
					}
					fmt.Println("ID:", t.Id, "Note title:", decTitle, "Description:", decMeta)
				}
			case entity.Card:
				cards := res.([]entity.CutCard)
				for _, c := range cards {
					numberDec, err := encrypt.DecryptMessage([]byte(key), c.CardNumber)
					if err != nil {
						fmt.Println(err)
						return
					}
					metaDec, err := encrypt.DecryptMessage([]byte(key), c.Meta)
					if err != nil {
						fmt.Println(err)
						return
					}
					fmt.Println("ID:", c.Id, "Card number:", numberDec, "Description:", metaDec)
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
