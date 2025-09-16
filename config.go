package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"strconv"
)

type Config struct {
	LogLevel slog.Level
	Port     uint
}

func (c *Config) Validate() error {
	if c.Port <= 0 {
		return fmt.Errorf("invalid port %d, expected to be > 0", c.Port)
	}

	return nil
}

func loadConfig() (*Config, error) {
	// defaults / optional config values
	conf := &Config{
		LogLevel: slog.LevelInfo,
		Port:     8080,
	}

	var err error

	err = setEnvConfigVals(conf)
	if err != nil {
		return nil, err
	}

	// cli flags has greatest precedence so is set last
	err = setFlagsConfigVals(conf)
	if err != nil {
		return nil, err
	}

	return conf, nil
}

// pulls config from env variables
// mutates provided config, only overriding for existing env variables
func setEnvConfigVals(conf *Config) error {
	if logLevelStr, found := lookupStrEnv("LOG_LEVEL"); found {
		logLevel, err := strToLogLevel(logLevelStr)
		if err != nil {
			return fmt.Errorf("failed to parse LOG_LEVEL env var: %w", err)
		}
		conf.LogLevel = logLevel
	}
	if port, found := lookupUintEnv("PORT"); found {
		conf.Port = port
	}

	return nil
}

// pulls config from cli flags
// mutates provided config, only overriding for existing cli flags
func setFlagsConfigVals(conf *Config) error {
	logLevelPtr := flag.String("log", "", "a log level string")
	portPtr := flag.Uint("port", 0, "a positive int > 0")

	flag.Parse()

	var lastErr error

	// using `Visit` to only set flags that were actually provided
	flag.Visit(func(f *flag.Flag) {
		switch f.Name {

		case "log":
			logLevel, err := strToLogLevel(*logLevelPtr)
			if err != nil {
				lastErr = fmt.Errorf("failed to parse -log flag: %w", err)
				return
			}
			conf.LogLevel = logLevel

		case "port":
			conf.Port = *portPtr

		}
	})

	return lastErr
}

// returns bool indicating whether key was found
func lookupStrEnv(key string) (val string, found bool) {
	return os.LookupEnv(key)
}

// returns bool indicating whether key was found
func lookupUintEnv(key string) (val uint, found bool) {
	strVal, found := os.LookupEnv(key)

	if parsedVal, err := strconv.ParseUint(strVal, 10, 32); err != nil {
		return uint(parsedVal), found
	}

	return 0, found
}

// attempts to parse a string into a slog.Level
func strToLogLevel(str string) (slog.Level, error) {
	var level slog.Level
	if err := level.UnmarshalText([]byte(str)); err != nil {
		return level, fmt.Errorf("invalid log level %q: %w", str, err)
	}
	return level, nil
}
