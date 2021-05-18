package cmds

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/keftcha/ccv/format"
)

// Capitalize is cobra command to capitalize text
var Capitalize = &cobra.Command{
	Use:   "capitalize",
	Short: "Capitalize the text",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(
			format.Capitalize(
				strings.Join(args, " "),
			),
		)
	},
}
