package login

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	var login, password string
	loginCmd := &cobra.Command{
		Use:   "login",
		Short: "Login to armGOur service",
		Run: func(cmd *cobra.Command, args []string) {
			//c.sender.loginUser(c.login, c.password)
			fmt.Println(login, password)
			err := os.WriteFile(".token", []byte(login+password), 0644)
			if err != nil {
				fmt.Println("Error writing to file:", err)
				return
			}

			fmt.Println("String written to file successfully.")
		},
	}
	loginCmd.PersistentFlags().StringVarP(
		&login, "login", "l", "", "User login")
	err := loginCmd.MarkPersistentFlagRequired("login")
	if err != nil {
		// TODO
		return nil
	}
	loginCmd.PersistentFlags().StringVarP(
		&password, "password", "p", "", "User password")
	err = loginCmd.MarkPersistentFlagRequired("password")
	if err != nil {
		// TODO
		return nil
	}
	loginCmd.MarkFlagsRequiredTogether("login", "password")
	return loginCmd
}
