package credentials

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/Albitko/secrets-armgour/internal/utils/encrypt"
)

type sender interface {
	CreateCredentials(serviceName, serviceLogin, servicePassword, meta, user string) error
}

func New(s sender) *cobra.Command {
	var serviceName, serviceLogin, servicePassword, meta string
	createCmd := &cobra.Command{
		Use:   "credentials",
		Short: "Save user credentials",
		Run: func(cmd *cobra.Command, args []string) {
			key, user, err := encrypt.GetCliSecrets()
			encName, err := encrypt.EncryptMessage([]byte(key), serviceName)
			encLogin, err := encrypt.EncryptMessage([]byte(key), serviceLogin)
			encPass, err := encrypt.EncryptMessage([]byte(key), servicePassword)
			encMeta, err := encrypt.EncryptMessage([]byte(key), meta)
			err = s.CreateCredentials(encName, encLogin, encPass, encMeta, user)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("Credentials successfully written")
		},
	}
	createCmd.PersistentFlags().StringVarP(
		&serviceName, "service", "s", "", "Service name")
	err := createCmd.MarkPersistentFlagRequired("service")

	createCmd.PersistentFlags().StringVarP(
		&serviceLogin, "login", "l", "", "Service login")
	err = createCmd.MarkPersistentFlagRequired("login")
	createCmd.PersistentFlags().StringVarP(
		&servicePassword, "password", "p", "", "Service password")
	err = createCmd.MarkPersistentFlagRequired("password")
	createCmd.PersistentFlags().StringVarP(
		&meta, "meta", "m", "", "Additional info about secret")
	if err != nil {
		// TODO
		return nil
	}
	return createCmd
}
