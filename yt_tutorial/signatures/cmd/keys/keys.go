package keys

import (
	"github.com/spf13/cobra"
)

var keysCmd = &cobra.Command{
	Use:   "keys",
	Short: "Manage keys pairs.",
	Long:  `Use subcommands to create public/private key pairs in PEM files.`,
}

func Init(rootCmd *cobra.Command) {
	rootCmd.AddCommand(keysCmd)
}
