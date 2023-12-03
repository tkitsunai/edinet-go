package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
	"net/url"
)

var storeCmd = &cobra.Command{
	Use:   "store",
	Short: "short",
	Long:  "long",
	RunE: func(cmd *cobra.Command, args []string) error {
		return storeBatch()
	},
}

var (
	fromDate string
	toDate   string
)

func init() {
	rootCmd.AddCommand(storeCmd)
	storeCmd.Flags().StringVarP(&fromDate, "from", "f", "", "use YYYY-MM-DD")
	storeCmd.Flags().StringVarP(&toDate, "to", "t", "", "use YYYY-MM-DD")
	_ = storeCmd.MarkFlagRequired("from")
	_ = storeCmd.MarkFlagRequired("to")
}

func storeBatch() error {
	client := http.DefaultClient
	base := "http://localhost:3000/documents"
	u, _ := url.Parse(base)
	q := u.Query()
	q.Set("from", fromDate)
	q.Set("to", toDate)
	u.RawQuery = q.Encode()

	response, err := client.Post(u.String(), "application/json", nil)
	if err != nil {
		return err
	}
	fmt.Println(response)
	return nil
}
