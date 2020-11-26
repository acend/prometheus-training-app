package cmd

import (
	"bytes"
	"encoding/json"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"net/http"
)

var webhookCmd = &cobra.Command{
	Use:   "webhook",
	Short: "Start the webhook application",
	Long: `Start the webhook application for getting alerts sent by Alertmanager`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			b, err := ioutil.ReadAll(r.Body)
			if err != nil {
				panic(err)
			}
			defer r.Body.Close()
			var buf bytes.Buffer
			if err := json.Indent(&buf, b, " >", "  "); err != nil {
				panic(err)
			}
			log.Println(buf.String())
		})))

	},
}

func init() {
	rootCmd.AddCommand(webhookCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
