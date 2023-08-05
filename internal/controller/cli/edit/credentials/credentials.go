package credentials

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/Albitko/secrets-armgour/internal/utils/encrypt"
)

type sender interface {
	EditCredentials(index int, serviceName, serviceLogin, servicePassword, meta string) error
}

func New(s sender) *cobra.Command {
	var serviceName, serviceLogin, servicePassword, meta string
	var index int
	editCmd := &cobra.Command{
		Use:   "credentials",
		Short: "Edit user credentials",
		Run: func(cmd *cobra.Command, args []string) {
			key, _, err := encrypt.GetCliSecrets()
			encName, err := encrypt.EncryptMessage([]byte(key), serviceName)
			encLogin, err := encrypt.EncryptMessage([]byte(key), serviceLogin)
			encPass, err := encrypt.EncryptMessage([]byte(key), servicePassword)
			encMeta, err := encrypt.EncryptMessage([]byte(key), meta)
			err = s.EditCredentials(index, encName, encLogin, encPass, encMeta)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("Credentials updated successfully.")

		},
	}
	editCmd.PersistentFlags().StringVarP(
		&serviceName, "service", "s", "", "Service name")
	err := editCmd.MarkPersistentFlagRequired("service")

	editCmd.PersistentFlags().StringVarP(
		&serviceLogin, "login", "l", "", "Service login")
	err = editCmd.MarkPersistentFlagRequired("login")
	editCmd.PersistentFlags().StringVarP(
		&servicePassword, "password", "p", "", "Service password")
	err = editCmd.MarkPersistentFlagRequired("password")
	editCmd.PersistentFlags().StringVarP(
		&meta, "meta", "m", "", "Additional info about secret")
	if err != nil {
		// TODO
		return nil
	}
	editCmd.Flags().IntVarP(
		&index, "index", "i", 0, "index")
	return editCmd
}
