package config

import (
	"github.com/spf13/cobra"
)

type Flags struct {
	ConfigPath string
}

func ParseFlags(cmd *cobra.Command) func() *Flags {
	return func() *Flags {
		configPath, err := cmd.Flags().GetString("config")
		if err != nil {
			return nil
		}

		return &Flags{
			ConfigPath: configPath,
		}
	}
}
