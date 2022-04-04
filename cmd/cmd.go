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

	// New variable of type *viper.Viper that stores and merges file configs (includes env vars)
	v := viper.New()

	// Base command of actual program
	rootCmd := &cobra.Command{
		Use:   "cli_args",
		Short: "Program to test cli args",
		Long:  `I literally have no idea what I'm doing, just go along with it`,
		Run: func(cmd *cobra.Command, args []string) {

			// On each rootCmd flag, if the flag is not "config", bind it to viper equivalent key
			cmd.Flags().VisitAll(func(f *pflag.Flag) {
				if f.Name != "config" {

					// Create an env var key named the same as flag, e.g. "foo-bar". The env var key is for accessing by code.
					// Take same env var key name, and normalize it to env var naming specification, e.g. "FOO_BAR",
					// so when assigning FOO_BAR=baz, it maps to foo-bar
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

	// Here, we start defining a load of flags

	// Config is a special case. We only want it to be configurable from the command line, not from other configs.
	cfgFilePath := *rootCmd.Flags().StringP("config", "c", cfgPathNoExt+".json", "Path to a config file. Supported types are {json,yaml}")

	// These are flags that can be changed from configs
	rootCmd.Flags().BoolVarP(&cfg.UseColor, "use-color", "u", false, "Display colorized ouput")
	rootCmd.Flags().BoolVarP(&cfg.ShowFunny, "show-funny", "s", false, "Show the funny thing :D")
	rootCmd.Flags().StringVar(&cfg.CoolString, "cool-string", "bruh", "A cool string to show")

	rootCmd.Execute()

	v.SetConfigFile(cfgFilePath) // Config file path to load
	v.SetEnvPrefix(envPrefix)    // Prefix for all environment variables
	v.ReadInConfig()             // Load config file

}
