package config

import (
	"fmt"
	"github.com/spf13/viper"
)

// Config holds the application configuration values.
type Config struct {
	Port        string // Port on which the server will run
	DatabaseURL string // URL for the database connection
	Environment string // Application environment (e.g., development, production)
}

// LoadConfig initializes and loads the configuration from a file or environment variables.
// It prioritizes environment variables over file values, ensuring flexibility for deployment.
func LoadConfig() (*Config, error) {
	// Create a new instance of Viper to avoid using the global instance.
	v := viper.New()

	// Configure Viper to read from the .env file
	v.AddConfigPath("../../") // Specify the directory to look for the configuration file
	v.SetConfigName(".env")   // Specify the configuration file name (without extension)
	v.SetConfigType("env")    // Specify the configuration file format

	// Set default values for configuration keys
	setDefaults(v)

	// Read the configuration file
	if err := v.ReadInConfig(); err != nil {
		// Return an error if the configuration file cannot be read
		return nil, fmt.Errorf("failed to load config file: %w", err)
	}

	// Automatically bind environment variables to override file values
	v.AutomaticEnv()

	// Map the loaded configuration to the Config struct
	config, err := mapConfig(v)
	if err != nil {
		// Return an error if the mapping fails
		return nil, fmt.Errorf("failed to map configuration: %v", err)
	}

	return config, nil
}

// setDefaults sets default values for configuration keys.
// These values will be used if not specified in the configuration file or environment variables.
func setDefaults(v *viper.Viper) {
	v.SetDefault("PORT", "8080")               // Default port for the server
	v.SetDefault("ENVIRONMENT", "development") // Default application environment
}

// mapConfig maps the configuration values from Viper to the Config struct.
// It ensures type safety and provides a structured representation of the configuration.
func mapConfig(v *viper.Viper) (*Config, error) {
	return &Config{
		Port:        v.GetString("PORT"),         // Get the server port
		DatabaseURL: v.GetString("DATABASE_URL"), // Get the database connection
		Environment: v.GetString("ENVIRONMENT"),  // Get the application environment
	}, nil
}
