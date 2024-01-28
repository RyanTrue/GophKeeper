package commands

import (
	"fmt"
	"github.com/RyanTrue/GophKeeper/internal/models"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
)

func init() {
	rootCmd.AddCommand(getCredsCmd)
}

var getCredsCmd = &cobra.Command{
	Use:       "get-creds",
	Short:     "get a specific credentials using UID",
	Args:      cobra.MinimumNArgs(1),
	ValidArgs: []string{"uid"},
	Example:   `get-creds 1234`,
	Run: func(cmd *cobra.Command, args []string) {
		authorized, err := dependencies.Services.Auth.CheckAuthorized(cmd.Context())
		if err != nil {
			authError(cmd, err)
			if err := dependencies.Services.User.Delete(cmd.Context()); err != nil {
				log.Error().Err(err).Msg("Deleting user info")
			}
			return
		}
		if !authorized {
			unauthorized(cmd)
			return
		}

		uidString := args[0]

		uid, err := strconv.Atoi(uidString)
		if err != nil {
			cmd.PrintErrf("Unable to convert [%s] to integer", uidString)
			return
		}

		syncCreds(cmd)

		secret, err := dependencies.Services.CredsSecret.Get(cmd.Context(), int64(uid))
		if err != nil {
			log.Error().Err(err).Msg("Getting creds from command")
			return
		}

		displayCredsSecret(cmd, secret)
	},
}

func displayCredsSecret(cmd *cobra.Command, secret *models.CredsSecret) {
	if secret == nil {
		cmd.PrintErrln("No secret with this ID found")
		return
	}

	line := strings.Repeat("-", 10)
	cmd.Println(fmt.Sprintf("%s %s %s", line, "Credentials", line))

	cmd.Printf("UID: %d\n", secret.UID)
	cmd.Printf("Website: %s\n", secret.Website)
	cmd.Printf("Login: %s\n", secret.Login)
	cmd.Printf("Password: %s\n", secret.Password)
	cmd.Printf("Additional data: %s\n", secret.AdditionalData)
}
