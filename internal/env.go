package internal

import (
	"fmt"
	"os"
)

var (
	flagEnv Environment
)

type Environment string

const (
	EnvDevelopment Environment = "dev"
	EnvStaging     Environment = "staging"
	EnvProduction  Environment = "production"
)

func init() {
	env := os.Getenv("ENV")
	switch env {
	case string(EnvDevelopment), "":
		flagEnv = EnvDevelopment
	case string(EnvStaging):
		flagEnv = EnvStaging
	case string(EnvProduction):
		flagEnv = EnvProduction
	default:
		panic(fmt.Sprintf("invalid ENV: %q", env))
	}
}

func IsDev() bool {
	return flagEnv == EnvDevelopment
}

func IsStaging() bool {
	return flagEnv == EnvStaging
}

func IsProduction() bool {
	return flagEnv == EnvProduction
}
