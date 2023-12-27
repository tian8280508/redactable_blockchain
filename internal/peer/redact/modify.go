package redact

import (
	"github.com/spf13/cobra"
)

var redactModifyCmd = &cobra.Command{
	Use:   "modify",
	Short: "modify a tx, todo",
	Long:  "todo",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func modifyCmd() *cobra.Command {
	return redactModifyCmd
}
