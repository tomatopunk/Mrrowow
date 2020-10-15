package table

import (
	"fmt"
	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
)

//todo new Table
func NewCmd() *cobra.Command{
	cmd := &cobra.Command{
		Use: "new",
		Args: cobra.ExactArgs(0),
		Short: "new table",
		Long: heredoc.Doc("new table"),
		Example: heredoc.Doc(""),
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("new table ~~~")
			return nil
		},
	}
	return cmd
}