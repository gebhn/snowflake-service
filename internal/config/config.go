package config

import (
	"fmt"
	"os"
)

func GetTursoDbUrl() string {
	return readEnvVar("TURSO_DB_URL", "libsql://my-super-db.turso.io")
}

func GetTursoDbToken() string {
	return readEnvVar("TURSO_DB_TOKEN", "super.secret_token")
}

func readEnvVar(envVar, suggestion string) string {
	if value, ok := os.LookupEnv(envVar); ok {
		return value
	}
	panic(fmt.Sprintf("env var %s is not set, suggested value: %s", envVar, suggestion))
}
