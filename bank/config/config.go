package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

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

type confVars struct {
	missing   []string // names of the mandatory environment variables that are missing
	malformed []string // errors describing malformed environment variable values
}

type EnvReader interface {
    Getenv(key string) string
}

type OsEnvReader struct{}

func (OsEnvReader) Getenv(key string) string {
    return os.Getenv(key)
}

func New(envReader EnvReader) (*Config, error) {
	vars := &confVars{}

	postgresHost := vars.mandatory(envReader, "POSTGRES_HOST")
	postgresPort := vars.mandatory(envReader, "POSTGRES_PORT")
	postgresUser := vars.mandatory(envReader, "POSTGRES_USER")
	postgresDB := vars.mandatory(envReader, "POSTGRES_DB")
	postgresPassword := vars.mandatory(envReader, "POSTGRES_PASSWORD")

	postgresSSLMode := vars.optional(envReader, "POSTGRES_SSL_MODE", "disable")
	postgresRootCertLoc := vars.optional(envReader, "POSTGRES_ROOT_CERT_LOC", "")

	postgresMaxOpenConns := vars.optionalInt(envReader, "POSTGRES_MAX_OPEN_CONNS", 10)
	postgresMaxIdleConns := vars.optionalInt(envReader, "POSTGRES_MAX_IDLE_CONNS", 5)
	postgresMaxIdleTime := vars.optionalDuration(envReader, "POSTGRES_MAX_IDLE_TIME", 5*time.Minute)

	if err := vars.Error(); err != nil {
		return nil, fmt.Errorf("error loading configuration: %w", err)
	}

	return &Config{
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
	}, nil
}

func (vars *confVars) optional(envReader EnvReader, key, fallback string) string {
    value := envReader.Getenv(key)
    if value == "" {
        return fallback
    }
    return value
}

func (vars *confVars) optionalInt(envReader EnvReader, key string, fallback int) int {
	value := envReader.Getenv(key)
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

func (vars *confVars) optionalDuration(envReader EnvReader, key string, fallback time.Duration) time.Duration {
	value := envReader.Getenv(key)
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

func (vars *confVars) mandatory(envReader EnvReader, key string) string {
	value := envReader.Getenv(key)
	if value == "" {
		vars.missing = append(vars.missing, key)
	}
	return value
}

func (vars confVars) Error() error {
	if len(vars.missing) > 0 {
		return fmt.Errorf("missing mandatory configurations: %s", strings.Join(vars.missing, ", "))
	}

	if len(vars.malformed) > 0 {
		return fmt.Errorf("malformed configurations: %s", strings.Join(vars.malformed, "; "))
	}
	return nil
}
