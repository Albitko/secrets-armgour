package credentials

import (
	"fmt"

	"github.com/spf13/cobra"
)

type sender interface {
	CreateCredentials(serviceName, serviceLogin, servicePassword, meta string) error
}

func New(s sender) *cobra.Command {
	var serviceName, serviceLogin, servicePassword, meta string
	createCmd := &cobra.Command{
		Use:   "credentials",
		Short: "Save user credentials",
		Run: func(cmd *cobra.Command, args []string) {
			err := s.CreateCredentials(serviceName, serviceLogin, servicePassword, meta)
			if err != nil {
				fmt.Println(err)
			}
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
