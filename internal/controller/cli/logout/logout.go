package logout

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	logoutCmd := &cobra.Command{
		Use:   "logout",
		Short: "Logout from armGOur service",
		Run: func(cmd *cobra.Command, args []string) {
			//c.sender.Logout()
			err := os.Remove(".token")
			if err != nil {
				fmt.Println("Error removing file:", err)
				return
			}
			fmt.Println("File removed successfully.")
		},
	}
	return logoutCmd
}
