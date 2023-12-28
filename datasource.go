package dotenv

import (
	"os"
)

// EnvDataSource is a data source that loads data from the environment.
type EnvDataSource struct {
}

// Get returns the value of the given key.
func (EnvDataSource) Get(key string) any {
	return os.Getenv(key)
}

// // utils
// var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
// var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

// func toSnakeCase(str string) string {
// 	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
// 	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
// 	return snake
// }
