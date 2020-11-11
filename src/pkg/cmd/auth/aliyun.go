package auth

import (
	"fmt"
	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
)

type AlibabaCloudOptions struct {
	AccessKey string
	AccessSecret string
}
func NewAlibabaCloudAuthCmd() *cobra.Command{
	cmd := &cobra.Command{
		Use:   "aliyun",
		Short: "Login, logout, and refresh your authentication",
		Long: heredoc.Doc(`
			manage authentication.
		`),
	}

	cmd.AddCommand(newAlibabaCloudLoginCmd())
	cmd.AddCommand(newAlibabaCloudLogoutCmd())
	cmd.AddCommand(newAlibabaCloudShowCmd())
	return cmd
}

func newAlibabaCloudLoginCmd() *cobra.Command {
	opt := &AlibabaCloudOptions{}
	cmd := &cobra.Command{
		Use: "login",
		Args: cobra.ExactArgs(2),
		Short: "Authenticate with a Alibaba Cloud,Please Use `login <accessKey> <accessToken>`",
		Long: heredoc.Doc("Authenticate with a Alibaba Cloud"),
		Example: heredoc.Doc(""),
		RunE: func(cmd *cobra.Command, args []string) error {
			opt.AccessKey = args[0]
			opt.AccessSecret = args[1]
			return nil
		},
	}

	cmd.PersistentFlags().StringVarP(&opt.AccessKey,"key","k","","AccessKey")
	cmd.PersistentFlags().StringVarP(&opt.AccessSecret,"secret","s","","AccessSecret")

	cmd.MarkFlagRequired("key")
	cmd.MarkFlagRequired("secret")
	return cmd
}

func newAlibabaCloudShowCmd() *cobra.Command{
	//opt := &AlibabaCloudOptions{}
	cmd := &cobra.Command{
		Use: "show",
		Args: cobra.ExactArgs(0),
		Short: "Show your Authenticate Token For Alibaba Cloud ",
		Long: heredoc.Doc("Show Authenticate For Alibaba Cloud"),
		Example: heredoc.Doc(""),
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("Alibaba Cloud Access_Key,Access_Secret")
			//fmt.Println(cmd.PersistentFlags().Lookup("key").Value)
			//fmt.Println(cmd.PersistentFlags().Lookup("secret").Value)
			return nil
		},
	}
	return cmd
}

func newAlibabaCloudLogoutCmd() *cobra.Command{
	cmd := &cobra.Command{
		Use: "logout",
		Args: cobra.ExactArgs(0),
		Short: "Logout with Alibaba Cloud",
		Long: heredoc.Doc("Logout with Alibaba Cloud"),
		Example: heredoc.Doc(""),
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("Logout with Alibaba Cloud")
			return nil
		},
	}
	return cmd
}