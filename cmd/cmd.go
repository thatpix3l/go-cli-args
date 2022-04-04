package cmd

import (
	"fmt"
	"path"
	"strings"

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
	// Home of config files.
	// Neat and tidy according to freedesktop.org's base directory specifications.
	// ...And whatever windows does, I guess...
	cfgHomePath  = path.Join(xdg.ConfigHome, "cli_args")
	cfgPathNoExt = path.Join(cfgHomePath, cfgBasename)
)

// GenerateConfig loads and parses config files from different sources,
// parses them, and finally merges them together with a specific order precedence,
// storing them into the given *config.Config struct
func GenerateConfig(cfg *config.Config) {

	// Viper config
	v := viper.New()

	// Base command of actual program
	rootCmd := &cobra.Command{
		Use:   "cli_args",
		Short: "Program to test cli args",
		Long:  `I literally have no idea what I'm doing, just go along with it`,
		Run: func(cmd *cobra.Command, args []string) {

			// On each rootCmd flag, if it is not the config flag, bind it to viper equivalent key
			cmd.Flags().VisitAll(func(f *pflag.Flag) {
				if f.Name != "config" {

					// Bind current flag name to environment variable equivalent
					envKey := envPrefix + "_" + strings.ToUpper(strings.ReplaceAll(f.Name, "-", "_"))
					v.BindEnv(f.Name, envKey)

					// If current flag value has not been changed and config does have a value,
					// assign to flag the config value
					if !f.Changed && v.IsSet(f.Name) {
						flagVal := v.Get(f.Name)
						cmd.Flags().Set(f.Name, fmt.Sprintf("%v", flagVal))
					}

				}
			})
		},
	}
	// Here, we start defining a load of flags and stuff

	// The config file path used for loading, the default
	cfgFilePath := *rootCmd.Flags().StringP("config", "c", cfgPathNoExt+".json", "Path to a config file. Supported types are {json,yaml}")
	rootCmd.Flags().BoolVarP(&cfg.UseColor, "use-color", "u", false, "Display colorized ouput")
	rootCmd.Flags().BoolVarP(&cfg.ShowFunny, "show-funny", "s", false, "Show the funny thing :D")
	rootCmd.Flags().StringVar(&cfg.CoolString, "cool-string", "bruh", "A cool string to show")

	rootCmd.Execute()

	v.SetConfigFile(cfgFilePath)
	v.SetEnvPrefix(envPrefix)
	v.ReadInConfig()
	v.BindPFlags(pflag.CommandLine)

}
