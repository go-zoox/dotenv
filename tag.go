package dotenv

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

// DataSource defines the interface for loading data from a data source.
type DataSource interface {
	Get(key string) string
}

func decodeWithTagFromDataSource(ptr interface{}, tagName string, dataSource DataSource) error {
	t := reflect.TypeOf(ptr).Elem()
	v := reflect.ValueOf(ptr).Elem()

	for i := 0; i < t.NumField(); i++ {
		typ := t.Field(i)
		val := v.Field(i)

		kind := val.Kind()

		tagValueName := typ.Tag.Get(tagName)
		tabValueDfeault := typ.Tag.Get("default")
		tagValueRequired := typ.Tag.Get("required")
		tagValue := dataSource.Get(tagValueName)
		if tagValue == "" {
			tagValue = tabValueDfeault
		}

		if tagValueRequired == "true" && tagValue == "" {
			return fmt.Errorf("%s is required", tagValueName)
		}

		switch kind {
		case reflect.String:
			val.SetString(tagValue)
		case reflect.Bool:
			switch tagValue {
			case "true", "True", "TRUE":
				val.SetBool(true)
			case "false", "False", "FALSE":
				val.SetBool(false)
			default:
				return fmt.Errorf("%s is not bool", tagValueName)
			}
		case reflect.Int, reflect.Int64:
			v, err := strconv.ParseInt(tagValue, 10, 64)
			if err != nil {
				return fmt.Errorf("%s is not int", tagValueName)
			}
			val.SetInt(v)
		case reflect.Float64:
			v, err := strconv.ParseFloat(tagValue, 64)
			if err != nil {
				return fmt.Errorf("%s is not float", tagValueName)
			}
			val.SetFloat(v)
		}
	}

	return nil
}

// Decode decodes the given struct pointer from the environment.
func Decode(ptr interface{}) error {
	return decodeWithTagFromDataSource(ptr, "env", &Env{})
}

// Env is a data source that loads data from the environment.
type Env struct {
}

// Get returns the value of the given key.
func (Env) Get(key string) string {
	return os.Getenv(key)
}
