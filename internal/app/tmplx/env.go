package tmplx

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

// getEnvVarsWithPrefix returns all environment variables that start with the specified prefix.
func getEnvVarsWithPrefix(prefix string) map[string]string {
	envVars := make(map[string]string)

	// Iterate through all environment variables
	for _, env := range os.Environ() {
		// Split key and value
		pair := strings.SplitN(env, "=", 2)
		if len(pair) != 2 {
			continue
		}

		key := pair[0]
		value := pair[1]

		if strings.HasPrefix(key, prefix) {
			envVars[key] = value
		}
	}

	return envVars
}

// readEnvFileWithPrefix reads all environment variables that start with the specified prefix from the given filename and returns them.
func readEnvFileWithPrefix(filename string, prefix string) map[string]string {
	envMap := make(map[string]string)

	if len(filename) > 0 {
		if envVars, err := godotenv.Read(filename); err != nil {
			log.Printf("failed to load env: %s\n", err.Error())
		} else {
			for key, value := range envVars {
				if strings.HasPrefix(key, prefix) {
					envMap[key] = value
				}
			}
			log.Printf(".env file loaded from '%s'\n", filename)
		}
	}

	return envMap
}

// LoadEnv loads environment variables that start with the specified prefix and environment variables from the envFile.
// If an environment variable with the same name is defined, the environment variable value takes precedence over the envFile setting.
// If envFile is not provided, only environment variables are loaded.
func LoadEnv(envFile string, prefix string) map[string]string {
	envMap := make(map[string]string)

	// Load only if envFile exists
	if envFile != "" {
		if _, err := os.Stat(envFile); err == nil {
			envMap = readEnvFileWithPrefix(envFile, prefix)
		} else {
			log.Printf(".env file not found: %s\n", envFile)
		}
	}

	// Load environment variables (overrides envFile settings)
	for key, value := range getEnvVarsWithPrefix(prefix) {
		envMap[key] = value
	}

	return envMap
}
