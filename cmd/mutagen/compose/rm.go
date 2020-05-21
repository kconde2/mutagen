package compose

import (
	"fmt"

	"github.com/spf13/cobra"
)

func rmMain(_ *cobra.Command, arguments []string) {
	// Handle top-level help and version flags.
	handleTopLevelHelp()
	handleTopLevelVersion()

	// TODO: Implement.
	fmt.Println("rm not yet implemented")
}

var rmCommand = &cobra.Command{
	Use:          "rm",
	Run:          rmMain,
	SilenceUsage: true,
}

var rmConfiguration struct {
	// help indicates the presence of the -h/--help flag.
	help bool
}

func init() {
	// Avoid Cobra's built-in help functionality that's triggered when the
	// -h/--help flag is present. We still explicitly register a -h/--help flag
	// below for shell completion support.
	rmCommand.SetHelpFunc(commandHelp)

	// Grab a handle for the command line flags.
	flags := rmCommand.Flags()

	// Wire up rm command flags.
	flags.BoolVarP(&rmConfiguration.help, "help", "h", false, "")
	// TODO: Wire up remaining flags.
}
