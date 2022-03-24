package dotenv

import (
	"os"

	"github.com/joho/godotenv"
)

func Load(filenames ...string) error {
	return godotenv.Load(filenames...)
}

func Get(key string, defaultValue ...string) string {
	if len(defaultValue) > 1 {
		panic("Too many arguments supplied")
	}

	value := os.Getenv(key)
	if value == "" && len(defaultValue) > 0 {
		return defaultValue[0]
	}

	return value
}
