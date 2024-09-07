package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	DEBUG       bool
	ENVIRONMENT string
	SERVICES    []string
}

func LoadConfig() (Config, error) {
	config := Config{}
	envVars := map[string]interface{}{
		"ENVIRONMENT": &config.ENVIRONMENT,
		"DEBUG":       &config.DEBUG,
		"SERVICES":    &config.SERVICES,
	}

	for key, target := range envVars {
		if err := Process_envs(key, target); err != nil {
			return Config{}, err
		}
	}
	return config, nil
}

func Process_envs(key string, target interface{}) error {
	value, exists := os.LookupEnv(key)
	if !exists {
		return errors.New("Environment variable %s not set" + key)
	}
	switch v := target.(type) {
	case *string:
		*v = value
	case *int:
		parsedValue, err := strconv.Atoi(value)
		if err != nil {
			return fmt.Errorf("invalid value for environment variable %s: %s", key, err)
		}
		*v = parsedValue
	case *bool:
		parsedValue, err := strconv.ParseBool(value)
		if err != nil {
			return fmt.Errorf("invalid value for environment variable %s: %s", key, err)
		}
		*v = parsedValue
	case *[]string:
		*v = strings.Split(value, ",")
	default:
		return fmt.Errorf("invalid value for environment variable %s", key)
	}
	return nil
}
