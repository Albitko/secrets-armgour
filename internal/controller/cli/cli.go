package cli

import (
	"github.com/spf13/cobra"
)

type sender interface {
}

type cliCommands struct {
	s sender
}

func (c *cliCommands) Cmd() *cobra.Command {
	return &cobra.Command{}
}

func (c *cliCommands) Login() *cobra.Command {
	return &cobra.Command{}
}

func (c *cliCommands) Logout() *cobra.Command {
	return &cobra.Command{}
}
func (c *cliCommands) Register() *cobra.Command {
	return &cobra.Command{}
}
func (c *cliCommands) List() *cobra.Command {
	return &cobra.Command{}
}

func (c *cliCommands) Get() *cobra.Command {
	return &cobra.Command{}
}

func (c *cliCommands) Create() *cobra.Command {
	return &cobra.Command{}
}
func (c *cliCommands) Edit() *cobra.Command {
	return &cobra.Command{}
}
func (c *cliCommands) Delete() *cobra.Command {
	return &cobra.Command{}
}

func (c *cliCommands) GeneratePassword() *cobra.Command {
	return &cobra.Command{}
}

func New(s sender) *cliCommands {
	return &cliCommands{
		s: s,
	}
}
