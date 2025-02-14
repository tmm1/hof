package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/hofstadter-io/hof/cmd/hof/flags"

	"github.com/hofstadter-io/hof/lib/gen"
)

var genLong = `  generate all the things, from code to data to config...`

func init() {

	GenCmd.Flags().BoolVarP(&(flags.GenFlags.Stats), "stats", "s", false, "Print generator statistics")
	GenCmd.Flags().StringSliceVarP(&(flags.GenFlags.Generator), "generator", "g", nil, "Generators to run, default is all discovered")
}

func GenRun(args []string) (err error) {

	// you can safely comment this print out
	// fmt.Println("not implemented")

	err = gen.Gen(args, flags.GenFlags)

	return err
}

var GenCmd = &cobra.Command{

	Use: "gen [files...]",

	Aliases: []string{
		"G",
	},

	Short: "generate code, data, and config from your data models and designs",

	Long: genLong,

	PreRun: func(cmd *cobra.Command, args []string) {

	},

	Run: func(cmd *cobra.Command, args []string) {
		var err error

		// Argument Parsing

		err = GenRun(args)
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

	ohelp := GenCmd.HelpFunc()
	ousage := GenCmd.UsageFunc()
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

	GenCmd.SetHelpFunc(help)
	GenCmd.SetUsageFunc(usage)

}
