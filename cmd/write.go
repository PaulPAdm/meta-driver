package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/PaulPAdm/meta-driver/internal/adapters"
	"github.com/PaulPAdm/meta-driver/internal/core"
	"github.com/spf13/cobra"
)

var writeCommand = &cobra.Command{
	Use:   "write [file]",
	Short: "Write metadata to a file",
	Args:  cobra.ExactArgs(1), //?
	RunE: func(cmd *cobra.Command, args []string) error {
		filePath := args[0]

		dataFlag, err := cmd.Flags().GetString("data")
		if err != nil {
			return err
		}

		var metadata core.Metadata
		if err := json.Unmarshal([]byte(dataFlag), &metadata); err != nil {
			return err
		}

		storage := adapters.NewXattrAdapter()
		service := core.NewMetaService(storage)

		if err := service.Write(filePath, &metadata); err != nil {
			return fmt.Errorf("failed to write metadata: %w", err)
		}

		fmt.Println("write metadata successfully")
		return nil
	},
}

func init() {
	writeCommand.Flags().StringP("data", "d", "", "data file path")
	if err := writeCommand.MarkFlagRequired("data"); err != nil {
		return
	}
	rootCmd.AddCommand(writeCommand)
}
