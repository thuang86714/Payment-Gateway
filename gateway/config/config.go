package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Config struct holds all the configuration settings for the application
type Config struct {
	PostgresHost     string
	PostgresPort     string
	PostgresUser     string
	PostgresDB       string
	PostgresPassword string

	PostgresSSLMode      string
	PostgresRootCertLoc  string
	PostgresMaxOpenConns int
	PostgresMaxIdleConns int
	PostgresMaxIdleTime  time.Duration
}

// confVars struct is used to track missing or malformed configuration variables
type confVars struct {
	missing   []string // names of the mandatory environment variables that are missing
	malformed []string // errors describing malformed environment variable values
}

// Conf is a global variable to hold the application configuration
var Conf *Config

// New creates and returns a new Config instance, or an error if configuration loading fails
func New() (*Config, error) {
	vars := &confVars{}

	postgresHost := vars.mandatory("POSTGRES_HOST")
	postgresPort := vars.mandatory("POSTGRES_PORT")
	postgresUser := vars.mandatory("POSTGRES_USER")
	postgresDB := vars.mandatory("POSTGRES_DB")
	postgresPassword := vars.mandatory("POSTGRES_PASSWORD")

	postgresSSLMode := vars.optional("POSTGRES_SSL_MODE", "disable")
	postgresRootCertLoc := vars.optional("POSTGRES_ROOT_CERT_LOC", "")

	postgresMaxOpenConns := vars.optionalInt("POSTGRES_MAX_OPEN_CONNS", 10)
	postgresMaxIdleConns := vars.optionalInt("POSTGRES_MAX_IDLE_CONNS", 5)
	postgresMaxIdleTime := vars.optionalDuration("POSTGRES_MAX_IDLE_TIME", 5*time.Minute)

	if err := vars.Error(); err != nil {
		return nil, fmt.Errorf("error loading configuration: %w", err)
	}

	config := &Config{
		PostgresHost:     postgresHost,
		PostgresPort:     postgresPort,
		PostgresUser:     postgresUser,
		PostgresDB:       postgresDB,
		PostgresPassword: postgresPassword,
		PostgresSSLMode:      postgresSSLMode,
		PostgresRootCertLoc:  postgresRootCertLoc,
		PostgresMaxOpenConns: postgresMaxOpenConns,
		PostgresMaxIdleConns: postgresMaxIdleConns,
		PostgresMaxIdleTime:  postgresMaxIdleTime,
	}

	Conf = config

	return config, nil
}

// optional retrieves an optional environment variable, returning a fallback value if not set
func (vars *confVars) optional(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

// optionalInt retrieves an optional integer environment variable, returning a fallback value if not set or malformed
func (vars *confVars) optionalInt(key string, fallback int) int {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	valueInt, err := strconv.Atoi(value)
	if err != nil {
		vars.malformed = append(vars.malformed, key)
		return fallback
	}

	return valueInt
}

// optionalDuration retrieves an optional duration environment variable, returning a fallback value if not set or malformed
func (vars *confVars) optionalDuration(key string, fallback time.Duration) time.Duration {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	valueDuration, err := time.ParseDuration(value)
	if err != nil {
		vars.malformed = append(vars.malformed, key)
		return fallback
	}

	return valueDuration
}

// mandatory retrieves a mandatory environment variable, adding it to the missing list if not set
func (vars *confVars) mandatory(key string) string {
	value := os.Getenv(key)
	if value == "" {
		vars.missing = append(vars.missing, key)
	}
	return value
}

// Error returns an error if any mandatory configurations are missing or if any configurations are malformed
func (vars confVars) Error() error {
	if len(vars.missing) > 0 {
		return fmt.Errorf("missing mandatory configurations: %s", strings.Join(vars.missing, ", "))
	}

	if len(vars.malformed) > 0 {
		return fmt.Errorf("malformed configurations: %s", strings.Join(vars.malformed, "; "))
	}
	return nil
}
