package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewVersion(name string, version string) *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: fmt.Sprintf("Print the version number of %s", name),
		Long:  fmt.Sprintf("All software has versions. This is %s's", name),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("%s v%s\n", name, version)
		},
	}
}
