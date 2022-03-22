package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/io4io/terra/utils"
	"go.uber.org/zap"
)

const DefaultConfigFile = "config.json"

// Init initialize the config package.
func Init() error {
	filename := utils.AbsPath(DefaultConfigFile)
	zap.S().Infow("read config file", "filename", filename)
	configFile, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("open config file failed: %w", err)
	}
	defer configFile.Close()

	cfg := &RootConfig{
		filename: filename,
	}

	if err := json.NewDecoder(configFile).Decode(cfg); err != nil {
		return fmt.Errorf("decode config file in JSON format failed: %w", err)
	}

	rootConfig = cfg // replace the global instance
	return nil
}
