package cmddatamodel

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/hofstadter-io/hof/cmd/hof/flags"
)

var checkpointLong = `create a snapshot of the data model`

func init() {

	CheckpointCmd.Flags().StringVarP(&(flags.Datamodel__CheckpointFlags.Bump), "bump", "b", "patch", "type of version bump in [major,minor,patch,<semver>]")
}

func CheckpointRun(args []string) (err error) {

	// you can safely comment this print out
	fmt.Println("not implemented")

	return err
}

var CheckpointCmd = &cobra.Command{

	Use: "checkpoint",

	Aliases: []string{
		"cp",
		"x",
	},

	Short: "create a snapshot of the data model",

	Long: checkpointLong,

	PreRun: func(cmd *cobra.Command, args []string) {

	},

	Run: func(cmd *cobra.Command, args []string) {
		var err error

		// Argument Parsing

		err = CheckpointRun(args)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	extra := func(cmd *cobra.Command) bool {

		return false
	}

	ohelp := CheckpointCmd.HelpFunc()
	ousage := CheckpointCmd.UsageFunc()
	help := func(cmd *cobra.Command, args []string) {
		if extra(cmd) {
			return
		}
		ohelp(cmd, args)
	}
	usage := func(cmd *cobra.Command) error {
		if extra(cmd) {
			return nil
		}
		return ousage(cmd)
	}

	CheckpointCmd.SetHelpFunc(help)
	CheckpointCmd.SetUsageFunc(usage)

}
