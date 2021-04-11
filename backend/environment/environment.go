package environment

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

const (
	PROD  Env = "production"
	DEV   Env = "development"
	TEST  Env = "test"
	LOCAL Env = "local"
	EMPTY Env = ""
)

type Env string

func (e Env) String() string {
	return string(e)
}

func ReadEnv(environment Env) {

	env := Env(strings.ToLower(environment.String()))
	switch env {
	case PROD, DEV, TEST, LOCAL, EMPTY:
		godotenv.Load(".env." + env.String() + ".local")
		godotenv.Load(".env." + env.String())
		godotenv.Load()
	default:
		panic("unexpected environment " + env)
	}
}

func Getenv(key string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return ""
}

var (
	DATABASE  string = Getenv("DATABASE")
	SECRETKEY string = Getenv("SECRETKEY")
)
