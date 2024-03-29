package dotenv

import (
	"os"

	"github.com/go-zoox/fs"
	"github.com/go-zoox/tag"
	"github.com/joho/godotenv"
)

func init() {
	if fs.IsExist(fs.JoinPath(fs.CurrentDir(), ".env")) {
		LoadToEnv()
	}
}

// Load loads the .env file into the environment.
func Load(value interface{}, filenames ...string) error {
	if err := LoadToEnv(filenames...); err != nil {
		return err
	}

	tg := tag.New("env", &EnvDataSource{})
	if err := tg.Decode(value); err != nil {
		return err
	}

	return nil
}

// LoadToEnv loads the .env file into environment.
func LoadToEnv(filenames ...string) error {
	return godotenv.Load(filenames...)
}

// Get gets the value of the given key from system environment.
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
