package table

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
)

func NewTableCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "table <command>",
		Short: "new,start,stop,remove your table",
		Long: heredoc.Doc(``),
	}

	cmd.AddCommand(NewCmd())
	cmd.AddCommand(ListCmd())
	cmd.AddCommand(DetailCmd())
	cmd.AddCommand(RemoveCmd())
	return cmd
}