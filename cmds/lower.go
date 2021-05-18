package cmds

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// Lower is cobra command to lower case
var Lower = &cobra.Command{
	Use:   "lower",
	Short: "Convert to lower case",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(strings.ToLower(strings.Join(args, " ")))
	},
}
