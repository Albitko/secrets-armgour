package register

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	var login, password string

	registerCmd := &cobra.Command{
		Use:   "register",
		Short: "Register in armGOur service",
		Run: func(cmd *cobra.Command, args []string) {
			//c.sender.registerUser(c.login, c.password)
			fmt.Println(login, password)
			err := os.WriteFile(".token", []byte(login+password), 0644)
			if err != nil {
				fmt.Println("Error writing to file:", err)
				return
			}

			fmt.Println("String written to file successfully.")
		},
	}
	registerCmd.PersistentFlags().StringVarP(
		&login, "login", "l", "", "User login")
	err := registerCmd.MarkPersistentFlagRequired("login")
	if err != nil {
		// TODO
		return nil
	}
	registerCmd.PersistentFlags().StringVarP(
		&password, "password", "p", "", "User password")
	err = registerCmd.MarkPersistentFlagRequired("password")
	if err != nil {
		// TODO
		return nil
	}
	registerCmd.MarkFlagsRequiredTogether("login", "password")
	return registerCmd
}
