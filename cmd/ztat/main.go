package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ztock/ztat/internal/config"
	"github.com/ztock/ztat/pkg/logger"
)

var (
	name    = "ztat"
	version = "1.0.0"
)

var cfg *config.Config
var cfgFile string

var rootCmd = &cobra.Command{
	Use:     name,
	Version: version,
	Short:   "pharmacy statistical system",
	Long: `A command line tool to display pharmacy statistical data.
Complete documentation is available at https://github.com/ztock/ztat`,
	SilenceUsage: true,
	Args:         cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		// ctx, cancel := context.WithCancel(context.Background())
		// defer cancel()

		// Init logger
		logger := initLog(cfg)
		logger.Debugf("Load config success: %#v", cfg)

		return nil
	},
}

// Execute is the entry point of the command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logrus.Debugf("Execute error: %#v", err)
	}
}

func init() {
	// Init default config
	cfg := config.New()

	// Initialize cobra
	cobra.OnInitialize(initConfig)

	// Add flags
	flagSet := rootCmd.PersistentFlags()
	flagSet.StringVar(&cfg.Server.Addr, "server", cfg.Server.Addr, "set the address for server")
	flagSet.StringVar(&cfg.Metrics.Addr, "metrics", cfg.Metrics.Addr, "set the address for metrics server")
	flagSet.BoolVar(&cfg.Console, "console", false, "whether logger output records to the stdout")

	if err := viper.BindPFlags(rootCmd.PersistentFlags()); err != nil {
		panic(err)
	}
}

// initConfig reads in config file and ENV variables if set
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		cfgPath := fmt.Sprintf("%s/.%s", home, name)
		viper.AddConfigPath(cfgPath)
	}

	viper.SetEnvPrefix(name)
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		log.Fatalf("Using config file: %s", viper.ConfigFileUsed())
	}

	// Unmarshal config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf(errors.Wrap(err, "cannot unmarshal config").Error())
	}
}

func initLog(cfg *config.Config) logger.Logger {
	if cfg.Console == true {
		cfg.Logger.FilePath = ""
	}

	return logger.New(logger.Config{
		Level:    cfg.Logger.Level,
		FilePath: cfg.Logger.FilePath,
	})
}

func main() {
	Execute()
}
