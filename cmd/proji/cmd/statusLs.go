package cmd

import (
	"os"

	"github.com/nikoksr/proji/pkg/proji/storage"

	"github.com/jedib0t/go-pretty/table"
	"github.com/spf13/cobra"
)

var statusLsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List statuses",
	RunE: func(cmd *cobra.Command, args []string) error {
		return listStatuses(projiEnv.Svc)
	},
}

func init() {
	statusCmd.AddCommand(statusLsCmd)
}

func listStatuses(svc storage.Service) error {
	statuses, err := svc.LoadAllStatuses()
	if err != nil {
		return err
	}

	// Table header
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"ID", "Title", "Comment"})

	// Fill table
	for _, status := range statuses {
		t.AppendRow([]interface{}{status.ID, status.Title, status.Comment})
	}

	// Print the table
	t.Render()
	return nil
}
