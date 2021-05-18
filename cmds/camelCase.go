package cmds

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/keftcha/ccv/format"
)

// CamelCase is cobra command to camel case
var CamelCase = &cobra.Command{
	Use:   "camel",
	Short: "Convert to camel case",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(
			format.CamelCase(
				strings.Join(args, " "),
			),
		)
	},
}
