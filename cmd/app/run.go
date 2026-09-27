package app

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var Run = &cobra.Command{
	Use:   "run",
	Short: "run app",
	Run: func(cmd *cobra.Command, args []string) {
		log.Info().Msgf("run %s app!", args[0])
	},
}
