package table

import (
	"fmt"
	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
)

//todo table info
func DetailCmd() *cobra.Command{
	cmd := &cobra.Command{
		Use: "detail",
		Args: cobra.ExactArgs(0),
		Short: "write table info",
		Long: heredoc.Doc("write table info"),
		Example: heredoc.Doc(""),
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("table 1 active | Start:2020/10/1 22:00:00 | End:2020/10/2 2020:00:00")
			return nil
		},
	}
	return cmd
}