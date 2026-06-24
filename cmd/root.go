package cmd

import (
	"github.com/spf13/cobra"
)

const rootLongDesc = `Meta Driver is a cross-platform utility for reading and writing
file metadata according to the Meta Protocol standard.

It stores metadata using native OS mechanisms:
- Linux/macOS: extended attributes (xattr)
- Windows: NTFS Alternate Data Streams (ADS)`

var rootCmd = &cobra.Command{
	Use:     "meta-driver",
	Aliases: []string{"mdr"},
	Short:   "Meta Driver — read and write file metadata by the Meta Protocol",
	Long:    rootLongDesc,
}

func Execute() error {
	return rootCmd.Execute()
}
