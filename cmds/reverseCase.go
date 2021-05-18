package cmds

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/keftcha/ccv/format"
)

// ReverseCase is cobra command to reverse the case
var ReverseCase = &cobra.Command{
	Use:   "rev",
	Short: "Reverse text case",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(
			format.ReverseCase(
				strings.Join(args, " "),
			),
		)
	},
}
