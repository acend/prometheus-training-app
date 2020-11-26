package cmd

import (
	"prometheus-training-app/pkg/sampleapp"

	"github.com/spf13/cobra"
)

var port *int

var sampleAppCmd = &cobra.Command{
	Use:   "sampleapp",
	Short: "Start the sample application",
	Long: `Start the sample application`,
	Run: func(cmd *cobra.Command, args []string) {
		sampleapp.HttpServer(*port)
	},
}

func init() {
	rootCmd.AddCommand(sampleAppCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sampleAppCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	port = sampleAppCmd.Flags().Int("port", 8080, "port of sample app (default: 8080)")
}
