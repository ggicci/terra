package internal

import (
	"fmt"
	"os"
)

type Environment string

const (
	EnvDevelopment Environment = "dev"
	EnvStaging     Environment = "stag"
	EnvProduction  Environment = "prod"
)

var (
	flagEnv Environment
)

func init() {
	env := os.Getenv("IO4_TERRA_ENV")
	switch env {
	case string(EnvDevelopment):
		flagEnv = EnvDevelopment
	case string(EnvStaging):
		flagEnv = EnvStaging
	case string(EnvProduction):
		flagEnv = EnvProduction
	default:
		panic(fmt.Sprintf("invalid IO4_TERRA_ENV: %q", env))
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
