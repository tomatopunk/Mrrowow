package auth

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
)

func NewAuthCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "auth <command>",
		Short: "Login, logout, and refresh your authentication",
		Long: heredoc.Doc(`
			manage authentication.
		`),
	}

	cmd.AddCommand(NewAliyunAuthCmd())
	return cmd
}
