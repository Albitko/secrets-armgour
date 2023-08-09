package logout

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// New - return command for logout
func New() *cobra.Command {
	logoutCmd := &cobra.Command{
		Use:   "logout",
		Short: "Logout from armGOur service",
		Run: func(cmd *cobra.Command, args []string) {
			err := os.Remove(".token")
			if err != nil {
				fmt.Println("Error removing file:", err)
				return
			}
			fmt.Println("Logout successfully.")
		},
	}
	return logoutCmd
}
