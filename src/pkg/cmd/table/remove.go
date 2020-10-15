package table

import (
	"fmt"
	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
)

//todo remove table
func RemoveCmd() *cobra.Command{
	cmd := &cobra.Command{
		Use: "remove",
		Args: cobra.ExactArgs(0),
		Short: "freed table",
		Long: heredoc.Doc("freed table"),
		Example: heredoc.Doc(""),
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("table 1 is freed")
			return nil
		},
	}
	return cmd
}