package table

import (
	"fmt"
	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
)

//todo table list
func ListCmd() *cobra.Command{
	cmd := &cobra.Command{
		Use: "list",
		Args: cobra.ExactArgs(0),
		Short: "write list table",
		Long: heredoc.Doc("write list table"),
		Example: heredoc.Doc(""),
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("table 1 active \n table2 active")
			return nil
		},
	}
	return cmd
}