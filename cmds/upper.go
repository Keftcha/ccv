package cmds

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// Upper is cobra command to upper case
var Upper = &cobra.Command{
	Use:   "upper",
	Short: "Convert to upper case",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(strings.ToUpper(strings.Join(args, " ")))
	},
}
