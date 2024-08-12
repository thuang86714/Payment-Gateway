package config

import (
    "os"
    "testing"
    "time"
)

func TestNew(t *testing.T) {
    // Test case 1: All mandatory fields are set
    t.Run("AllMandatoryFieldsSet", func(t *testing.T) {
        os.Setenv("POSTGRES_HOST", "testhost")
        os.Setenv("POSTGRES_PORT", "5432")
        os.Setenv("POSTGRES_USER", "testuser")
        os.Setenv("POSTGRES_DB", "testdb")
        os.Setenv("POSTGRES_PASSWORD", "testpass")
        
        config, err := New()
        if err != nil {
            t.Fatalf("Expected no error, got %v", err)
        }
        
        if config.PostgresHost != "testhost" {
            t.Errorf("Expected PostgresHost to be 'testhost', got '%s'", config.PostgresHost)
        }
        // Add similar checks for other fields
    })

    // Test case 2: Missing mandatory field
    t.Run("MissingMandatoryField", func(t *testing.T) {
        os.Unsetenv("POSTGRES_HOST")
        
        _, err := New()
        if err == nil {
            t.Fatal("Expected error due to missing mandatory field, got nil")
        }
    })

    // Test case 3: Optional fields with custom values
    t.Run("OptionalFieldsCustomValues", func(t *testing.T) {
        os.Setenv("POSTGRES_HOST", "testhost")
        os.Setenv("POSTGRES_PORT", "5432")
        os.Setenv("POSTGRES_USER", "testuser")
        os.Setenv("POSTGRES_DB", "testdb")
        os.Setenv("POSTGRES_PASSWORD", "testpass")
        os.Setenv("POSTGRES_SSL_MODE", "require")
        os.Setenv("POSTGRES_MAX_OPEN_CONNS", "20")
        os.Setenv("POSTGRES_MAX_IDLE_TIME", "10m")
        
        config, err := New()
        if err != nil {
            t.Fatalf("Expected no error, got %v", err)
        }
        
        if config.PostgresSSLMode != "require" {
            t.Errorf("Expected PostgresSSLMode to be 'require', got '%s'", config.PostgresSSLMode)
        }
        if config.PostgresMaxOpenConns != 20 {
            t.Errorf("Expected PostgresMaxOpenConns to be 20, got %d", config.PostgresMaxOpenConns)
        }
        if config.PostgresMaxIdleTime != 10*time.Minute {
            t.Errorf("Expected PostgresMaxIdleTime to be 10 minutes, got %v", config.PostgresMaxIdleTime)
        }
    })

    // Test case 4: Malformed optional field
    t.Run("MalformedOptionalField", func(t *testing.T) {
        os.Setenv("POSTGRES_HOST", "testhost")
        os.Setenv("POSTGRES_PORT", "5432")
        os.Setenv("POSTGRES_USER", "testuser")
        os.Setenv("POSTGRES_DB", "testdb")
        os.Setenv("POSTGRES_PASSWORD", "testpass")
        os.Setenv("POSTGRES_MAX_OPEN_CONNS", "not_a_number")
        
        _, err := New()
        if err == nil {
            t.Fatal("Expected error due to malformed optional field, got nil")
        }
    })
}

func TestOptional(t *testing.T) {
    vars := &confVars{}
    
    t.Run("UsesFallbackValue", func(t *testing.T) {
        os.Unsetenv("TEST_KEY")
        result := vars.optional("TEST_KEY", "fallback")
        if result != "fallback" {
            t.Errorf("Expected 'fallback', got '%s'", result)
        }
    })

    t.Run("UsesEnvironmentValue", func(t *testing.T) {
        os.Setenv("TEST_KEY", "envvalue")
        result := vars.optional("TEST_KEY", "fallback")
        if result != "envvalue" {
            t.Errorf("Expected 'envvalue', got '%s'", result)
        }
    })
}

// Add similar tests for optionalInt, optionalDuration, and mandatory methods

func TestError(t *testing.T) {
    t.Run("NoErrors", func(t *testing.T) {
        vars := &confVars{}
        err := vars.Error()
        if err != nil {
            t.Fatalf("Expected no error, got %v", err)
        }
    })

    t.Run("MissingMandatoryFields", func(t *testing.T) {
        vars := &confVars{missing: []string{"FIELD1", "FIELD2"}}
        err := vars.Error()
        if err == nil {
            t.Fatal("Expected error for missing mandatory fields, got nil")
        }
    })

    t.Run("MalformedFields", func(t *testing.T) {
        vars := &confVars{malformed: []string{"FIELD1", "FIELD2"}}
        err := vars.Error()
        if err == nil {
            t.Fatal("Expected error for malformed fields, got nil")
        }
    })
}