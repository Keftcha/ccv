package cmds

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/keftcha/ccv/format"
)

// AlternateFancy is cobra command to fancy the text
var AlternateFancy = &cobra.Command{
	Use:   "fancyAlt",
	Short: "Convert to alternative fancy text",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(format.AlternateFancy(strings.Join(args, " ")))
	},
}
