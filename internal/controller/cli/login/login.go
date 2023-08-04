package login

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/Albitko/secrets-armgour/internal/entity"
)

type sender interface {
	LoginUser(login, password string) error
}

func New(s sender) *cobra.Command {
	var login, password string
	loginCmd := &cobra.Command{
		Use:   "login",
		Short: "Login to armGOur service",
		Run: func(cmd *cobra.Command, args []string) {
			err := s.LoginUser(login, password)
			if errors.Is(err, entity.ErrInvalidCredentials) {
				fmt.Println("Invalid credentials")
				return
			}
			if err != nil {
				fmt.Println(err)
			}

			hasher := sha1.New()
			hasher.Write([]byte(login + password))
			encKey := hex.EncodeToString(hasher.Sum(nil))[:16]

			data := entity.CliSecrets{
				UserName: login,
				Token:    "TESTTOKEN",
				Key:      encKey,
			}
			jsonData, err := json.Marshal(data)
			if err != nil {
				fmt.Printf("could not marshal json: %s\n", err)
				return
			}
			err = os.WriteFile(".token", jsonData, 0644)
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
