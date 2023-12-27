package redact

import (
	"github.com/hyperledger/fabric/common/flogging"
	"github.com/spf13/cobra"
)

const (
	redactCmdName = "redact"
	redactCmdDesc = "Operate a peer redact: this is a test, todo"
)

var logger = flogging.MustGetLogger("tjz")

var redactCmd = &cobra.Command{
	Use:   redactCmdName,
	Short: redactCmdDesc,
	Long:  redactCmdDesc,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {

	},
}

// Cmd returns the cobra command for Redact
func Cmd() *cobra.Command {
	redactCmd.AddCommand(modifyCmd())
	return redactCmd
}
