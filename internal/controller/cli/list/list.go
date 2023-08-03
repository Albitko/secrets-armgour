package list

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/Albitko/secrets-armgour/internal/entity"
)

type sender interface {
	ListUserSecrets(data string) (interface{}, error)
}

func New(s sender) *cobra.Command {
	var data string
	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List user saved secrets",
		Run: func(cmd *cobra.Command, args []string) {
			res, err := s.ListUserSecrets(data)
			if err != nil {
				fmt.Println(err)
			}
			switch data {
			case "credentials":
				credentials := res.([]entity.CutCredentials)
				for _, c := range credentials {
					fmt.Println(c.Id, c.ServiceName, c.Meta)
				}
			case "binary":
				bin := res.([]entity.CutBinary)
				for _, b := range bin {
					fmt.Println(b.Id, b.Title, b.Meta)
				}
			case "text":
				texts := res.([]entity.CutText)
				for _, t := range texts {
					fmt.Println(t.Id, t.Title, t.Meta)
				}
			case "card":
				cards := res.([]entity.CutCard)
				for _, c := range cards {
					fmt.Println(c.Id, c.CardNumber, c.Meta)
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
