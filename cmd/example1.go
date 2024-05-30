package cmd

import (
	"github.com/gkwa/farmay/example1"
	"github.com/spf13/cobra"
)

// example1Cmd represents the example1 command
var example1Cmd = &cobra.Command{
	Use:   "example1",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		example1.Run()
	},
}

func init() {
	rootCmd.AddCommand(example1Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// example1Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only example1 when this command
	// is called directly, e.g.:
	// example1Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
