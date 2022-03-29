package dotenv

import (
	"os"

	"github.com/joho/godotenv"
)

func Load(value interface{}, filenames ...string) error {
	if err := godotenv.Load(filenames...); err != nil {
		return err
	}

	if err := Decode(value); err != nil {
		return err
	}

	return nil
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
