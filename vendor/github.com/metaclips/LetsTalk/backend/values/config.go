package values

import (
	"encoding/json"
	"os"
)

type configuration struct {
	TLS struct {
		CertPath string
		KeyPath  string
	}

	ICEServers []struct {
		URLs       []string
		Username   string
		AuthType   string
		AuthSecret string
	}

	CorsAllowedOrigins []string

	DbHost                   string
	DbName                   string
	DropboxToken             string
	EnableClassSessionRecord bool // If EnableClassSessionUpload is set to true and no token is provided, files are saved to DB using GridFS.
	Port                     string
}

// Config contains application environment variables.
var Config configuration

// LoadConfiguration loads all application environment variables.
func LoadConfiguration(configPath string) error {
	file, err := os.Open(configPath) // For read access.
	if err != nil {
		return err
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Config)
	if err != nil {
		return err
	}

	initIceServers()

	return nil
}
