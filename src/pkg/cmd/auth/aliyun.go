package auth

import (
	"fmt"
	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
)

func NewAliyunAuthCmd() *cobra.Command{
	cmd := &cobra.Command{
		Use:   "aliyun",
		Short: "Login, logout, and refresh your authentication",
		Long: heredoc.Doc(`
			manage authentication.
		`),
	}
	//Login,Logout
	cmd.AddCommand(newAliyunLoginCmd())
	cmd.AddCommand(newAliyunLogoutCmd())
	return cmd
}

//todo Authenticate with aliyun~~
func newAliyunLoginCmd() *cobra.Command{
	cmd := &cobra.Command{
		Use: "login",
		Args: cobra.ExactArgs(0),
		Short: "Authenticate with a aliyun",
		Long: heredoc.Doc("Authenticate with a aliyun"),
		Example: heredoc.Doc(""),
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("Authenticate with a aliyun~")
			return nil
		},
	}
	return cmd
}

func newAliyunLogoutCmd() *cobra.Command{
	cmd := &cobra.Command{
		Use: "logout",
		Args: cobra.ExactArgs(0),
		Short: "Logout with aliyun",
		Long: heredoc.Doc("Logout with aliyun"),
		Example: heredoc.Doc(""),
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("Logout with aliyun")
			return nil
		},
	}
	return cmd
}