package main

import (
	"os"
	"time"

	"github.com/eneco/landscaper/pkg/landscaper"
	"github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

var env = &landscaper.Environment{}

var (
	dryRun = false
)

var rootCmd = &cobra.Command{
	Use:   "landscaper",
	Short: "A landscape desired state applicator",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if env.Verbose {
			logrus.SetLevel(logrus.DebugLevel)
		}
		return nil
	},
	SilenceUsage: true,
}

func init() {
	_ = rootCmd.PersistentFlags()
	p := &prefixed.TextFormatter{
		ForceColors: true,
	}
	p.TimestampFormat = time.RFC3339
	logrus.SetFormatter(p)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
