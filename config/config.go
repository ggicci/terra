package config

var (
	rootConfig *RootConfig
)

type RootConfig struct {
	filename string

	Server  *ServerConfig
	TerraDB *Postgres
}

// GetConfig returns the root configuration object.
func GetConfig() *RootConfig { return rootConfig }

func replaceRootConfig(newConfig *RootConfig) {
	rootConfig = newConfig
}
