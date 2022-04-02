package cmd

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/adrg/xdg"
	"github.com/spf13/pflag"
	"github.com/thatpix3l/args/config"
)

// Load a config file when given valid string path.
func loadFileConfig(cfg_path string, cfg *config.Config) {

	jsonFile, _ := os.Open(cfg_path)
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, cfg)

}

// Merge modified values between a source config and destination config
func mergeConfigs(src_cfg *config.Config, dest_cfg *config.Config) {

	if cli_arg_bytes, err := json.Marshal(&src_cfg); err == nil {

		json.Unmarshal(cli_arg_bytes, &dest_cfg)

	} else {

		log.Fatal(err)

	}
}

func GenerateConfig() config.Config {

	// Process command-line arguments

	var cmd_args_cfg config.Config
	cmd_args_cfg.UseColor = pflag.Bool("colorize", false, "Display colorized output")
	cmd_args_cfg.ShowFunny = pflag.Bool("show-funny", false, "Show the funny thing :D")

	cfg_path := pflag.StringP("config", "c", path.Join(xdg.ConfigHome, "cli_args/config.json"), "Path to a valid JSON config file")

	pflag.Parse()

	// Load and merge different sources of configs

	var main_cfg config.Config
	loadFileConfig(*cfg_path, &main_cfg)
	mergeConfigs(&cmd_args_cfg, &main_cfg)

	return main_cfg

}
