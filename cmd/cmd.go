package cmd

import (
	"fmt"
	"path"

	"github.com/adrg/xdg"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/thatpix3l/args/config"
)

const (
	cfgBasename = "butter"
	envPrefix   = "BUTTER"
)

var (
	cfgDirPath            = path.Join(xdg.ConfigHome, "cli_args")
	cfgFilePathWithoutExt = path.Join(cfgDirPath, cfgBasename)
)

// GenerateConfig loads and parses config files from different sources,
// parses them, and finally merges them together with a specific order precedence,
// storing them into the given *config.Config struct
func GenerateConfig(cfg *config.Config) {

	rootCmd := &cobra.Command{
		Use:   "cli_args",
		Short: "Program to test cli args",
		Long:  `I literally have no idea what I'm doing, just go along with it`,
	}

	cfgFilePath := *rootCmd.Flags().StringP("config", "c", cfgFilePathWithoutExt, "Path to a config file. Supported types are {json,yaml}")
	rootCmd.Flags().BoolP("use-color", "u", false, "Display colorized ouput")
	rootCmd.Flags().BoolP("show-funny", "s", false, "Show the funny thing :D")
	rootCmd.Flags().String("cool-string", "bruh", "A cool string to show")

	rootCmd.Execute()

	v := viper.New()
	v.SetConfigFile(cfgFilePath)
	v.ReadInConfig()
	v.BindPFlags(pflag.CommandLine)

	fmt.Printf("%s\n", v.Get("use-color"))

}
