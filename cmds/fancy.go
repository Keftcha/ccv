package cmds

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/keftcha/ccv/format"
)

// Fancy is cobra command to fancy the text
var Fancy = &cobra.Command{
	Use:   "fancy",
	Short: "Convert to fancy text",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(format.Fancy(strings.Join(args, " ")))
	},
}
