// SPDX-License-Identifier: MIT
// Copyright (c) 2025 Geza Corp authors
package main

import (
	"context"
	"fmt"
	"os"

	"emperror.dev/errors"
	"github.com/gezacorp/licsensei/internal/licsensei"
	"github.com/go-logr/logr"
	"github.com/iand/logfmtr"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type HeaderConfig struct {
	Copyrights   []string
	Authors      []string
	LicenseTypes []string
	LicenseTexts []string
}

type Config struct {
	IgnorePaths   []string
	IgnoreFiles   []string
	HeaderConfigs []HeaderConfig
}

var (
	verbosity      int
	configFileName string
)

func init() {
	pflag.StringVarP(&configFileName, "config", "c", "", "config file path")
	pflag.IntVarP(&verbosity, "verbosity", "v", 0, "log verbosity, higher value produces more output")
	pflag.Parse()
}

func main() {
	opts := logfmtr.DefaultOptions()
	opts.Colorize = true
	opts.Humanize = true

	logger := logfmtr.NewWithOptions(opts)

	logfmtr.SetVerbosity(verbosity)

	ctx := logr.NewContext(context.Background(), logger)

	if configFileName != "" {
		viper.SetConfigFile(configFileName)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName(".licsensei")
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("can't read config:", err)
		os.Exit(1)
	}

	var conf Config
	if err := viper.GetViper().Unmarshal(&conf); err != nil {
		panic(errors.WrapIf(err, "failed to unmarshal configuration"))
	}

	root, err := os.Getwd()
	if err != nil {
		panic(errors.WrapIf(err, "could not get working directory"))
	}

	licenseCheckConfigs := []licsensei.LicenseCheckConfig{}
	for _, headerConfig := range conf.HeaderConfigs {
		licenseCheckConfigs = append(licenseCheckConfigs, licsensei.LicenseCheckConfig{
			Copyrights:   headerConfig.Copyrights,
			Authors:      headerConfig.Authors,
			LicenseTypes: headerConfig.LicenseTypes,
			LicenseTexts: headerConfig.LicenseTexts,
		})
	}

	results, err := licsensei.Check(ctx, root, licsensei.CheckConfiguration{
		IgnorePaths:         conf.IgnorePaths,
		IgnoreFiles:         conf.IgnoreFiles,
		LicenseCheckConfigs: licenseCheckConfigs,
	})
	if err != nil {
		panic(err)
	}

	for path, result := range results {
		logger.Error(errors.NewPlain(result), path)
	}

	if len(results) > 0 {
		os.Exit(1)
	}
}
