package distrprocess

import (
	"github.com/bosamatheus/distributed-processing/pkg/distrprocess"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "distrprocess",
	Short: "An example of distributed processing in Golang.",
	Run: func(cmd *cobra.Command, args []string) {
		res := distrprocess.Run()
		log.Info(res)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
