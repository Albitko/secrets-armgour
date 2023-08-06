package credentials

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/Albitko/secrets-armgour/internal/utils/encrypt"
)

type sender interface {
	CreateCredentials(serviceName, serviceLogin, servicePassword, meta, user string) error
}

// New - return command for credentials creation
func New(s sender) *cobra.Command {
	var serviceName, serviceLogin, servicePassword, meta string
	createCmd := &cobra.Command{
		Use:   "credentials",
		Short: "Save user credentials",
		Run: func(cmd *cobra.Command, args []string) {
			var key, user, encName, encLogin, encPass, encMeta string
			var err error
			key, user, err = encrypt.GetCliSecrets()
			if err != nil {
				fmt.Println("Error reading file:", err)
				return
			}
			encName, err = encrypt.EncryptMessage([]byte(key), serviceName)
			if err != nil {
				fmt.Println("Error reading file:", err)
				return
			}
			encLogin, err = encrypt.EncryptMessage([]byte(key), serviceLogin)
			if err != nil {
				fmt.Println("Error reading file:", err)
				return
			}
			encPass, err = encrypt.EncryptMessage([]byte(key), servicePassword)
			if err != nil {
				fmt.Println("Error reading file:", err)
				return
			}
			encMeta, err = encrypt.EncryptMessage([]byte(key), meta)
			if err != nil {
				fmt.Println("Error reading file:", err)
				return
			}
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
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}
	createCmd.PersistentFlags().StringVarP(
		&serviceLogin, "login", "l", "", "Service login")
	err = createCmd.MarkPersistentFlagRequired("login")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}
	createCmd.PersistentFlags().StringVarP(
		&servicePassword, "password", "p", "", "Service password")
	err = createCmd.MarkPersistentFlagRequired("password")
	createCmd.PersistentFlags().StringVarP(
		&meta, "meta", "m", "", "Additional info about secret")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return createCmd
}
