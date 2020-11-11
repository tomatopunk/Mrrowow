package cmdRoot

import (
	"catTable/src/pkg/cmd/auth"
	"catTable/src/pkg/cmd/io"
	"catTable/src/pkg/cmd/table"
	"catTable/src/pkg/cmd/version"
	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
)

func CreateCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "catt <command> <subcommand> [flags]",
		Short: "Cat Table",
		Long:  `Stop letting the cat near your table because it will push everything off your table!!!`,

		//SilenceErrors is an option to quiet errors down stream
		SilenceErrors: true,
		SilenceUsage:  true,
		Example: heredoc.Doc(`
			$ catt auth ali
			$ catt table new
			$ catt table stop
			$ catt table remove
			$ catt table list
		`),
		Annotations: map[string]string{
			"help:feedback": heredoc.Doc(`
				Open an issue 'https://github.com/tomatopunk/CatTable'
			`),
		},
	}

	io := io.GetIOStreams()

	cmd.SetOut(io.Out)
	cmd.SetErr(io.ErrOut)

	formatVersion := version.Format()

	cmd.Version = formatVersion
	cmd.SetVersionTemplate(formatVersion)

	// add child command
	cmd.AddCommand(auth.NewAuthCmd())
	cmd.AddCommand(table.NewTableCmd())

	return cmd
}
