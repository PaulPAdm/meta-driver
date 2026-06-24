package cmd

import (
	"fmt"

	"github.com/PaulPAdm/meta-driver/internal/adapters"
	"github.com/PaulPAdm/meta-driver/internal/core"
	"github.com/spf13/cobra"
)

var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Read data from a file",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		filePath := args[0]

		storage := adapters.NewXattrAdapter()
		service := core.NewMetaService(storage)

		data, err := service.Read(filePath)

		if err != nil {
			return fmt.Errorf("failed to read data: %w", err)
		}
		fmt.Println(*data)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(readCmd)
}
