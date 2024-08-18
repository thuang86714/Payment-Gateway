package config

import (
    "testing"
    "time"

    "github.com/stretchr/testify/assert"
)

type MockEnvReader struct {
    envVars map[string]string
}

func (m MockEnvReader) Getenv(key string) string {
    return m.envVars[key]
}

func TestNew(t *testing.T) {
    mockEnv := MockEnvReader{
        envVars: map[string]string{
            "POSTGRES_HOST":     "testhost",
            "POSTGRES_PORT":     "5432",
            "POSTGRES_USER":     "testuser",
            "POSTGRES_DB":       "testdb",
            "POSTGRES_PASSWORD": "testpass",
            "POSTGRES_MAX_OPEN_CONNS": "15",
            "POSTGRES_MAX_IDLE_TIME":  "10m",
        },
    }

    config, err := New(mockEnv)
    assert.NoError(t, err)
    assert.NotNil(t, config)

    assert.Equal(t, "testhost", config.PostgresHost)
    assert.Equal(t, "5432", config.PostgresPort)
    assert.Equal(t, "testuser", config.PostgresUser)
    assert.Equal(t, "testdb", config.PostgresDB)
    assert.Equal(t, "testpass", config.PostgresPassword)
    assert.Equal(t, "disable", config.PostgresSSLMode) // default value
    assert.Equal(t, 15, config.PostgresMaxOpenConns)
    assert.Equal(t, 5, config.PostgresMaxIdleConns) // default value
    assert.Equal(t, 10*time.Minute, config.PostgresMaxIdleTime)
}

func TestNew_MissingMandatory(t *testing.T) {
    mockEnv := MockEnvReader{
        envVars: map[string]string{},
    }

    _, err := New(mockEnv)
    assert.Error(t, err)
    assert.Contains(t, err.Error(), "missing mandatory configurations")
}

func TestNew_MalformedOptional(t *testing.T) {
    mockEnv := MockEnvReader{
        envVars: map[string]string{
            "POSTGRES_HOST":     "testhost",
            "POSTGRES_PORT":     "5432",
            "POSTGRES_USER":     "testuser",
            "POSTGRES_DB":       "testdb",
            "POSTGRES_PASSWORD": "testpass",
            "POSTGRES_MAX_OPEN_CONNS": "not_an_int",
        },
    }

    config, err := New(mockEnv)
    assert.NoError(t, err) // It should not error, but use the default value
    assert.Equal(t, 10, config.PostgresMaxOpenConns) // default value
}