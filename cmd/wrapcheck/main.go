package main

import (
	"log"

	"github.com/spf13/viper"
	"github.com/tomarrell/wrapcheck/v2/wrapcheck"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	viper.SetConfigName(".wrapcheck")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/.wrapcheck")
	viper.AddConfigPath(".")

	viper.SetDefault("ignoreSigs", wrapcheck.DefaultIgnoreSigs)

	// Read in config, ignore if the file isn't found and use defaults.
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			log.Fatalf("failed to parse config: %v", err)
		}
	}

	var cfg wrapcheck.WrapcheckConfig
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("failed to unmarshal config: %v", err)
	}

	singlechecker.Main(wrapcheck.NewAnalyzer(cfg))
}
