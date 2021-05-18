package cmds

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/keftcha/ccv/format"
)

// Titleize is cobra command to titleize text
var Titleize = &cobra.Command{
	Use:   "titleize",
	Short: "Titleize the text",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(
			format.Titleize(
				strings.Join(args, " "),
			),
		)
	},
}
