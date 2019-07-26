package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

// Load reads in config file and ENV variables if set.
func Load(cfgFile string) {
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

		// Search config in home directory with name ".zimt" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".zimt")
	}

	// flag "my-key" matches with "MY_KEY" environment variable
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AutomaticEnv() // read in environment variables that match

	if viper.GetBool("verbose") {
		fmt.Printf("using config file: %q\n", viper.ConfigFileUsed())
	}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("warning: config file is not found, using defaults")
			// Config file not found; ignore error if desired
		} else {
			log.Fatalf("error while reading config file: %s", err)
			// fmt.Fat
			// panic(fmt.Errorf("Fatal error config file: %s \n", err))
			// Config file was found but another error was produced
		}
	}
}
