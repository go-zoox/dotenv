package dotenv

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

type DataSource interface {
	Get(key string) string
}

func decodeWithTagFromDataSource(ptr interface{}, tagName string, dataSource DataSource) error {
	t := reflect.TypeOf(ptr).Elem()
	v := reflect.ValueOf(ptr).Elem()

	for i := 0; i < t.NumField(); i++ {
		type_ := t.Field(i)
		value_ := v.Field(i)

		kind := value_.Kind()

		tagValueName := type_.Tag.Get(tagName)
		tabValueDfeault := type_.Tag.Get("default")
		tagValueRequired := type_.Tag.Get("required")
		tagValue := dataSource.Get(tagValueName)
		if tagValue == "" {
			tagValue = tabValueDfeault
		}

		if tagValueRequired == "true" && tagValue == "" {
			return fmt.Errorf("%s is required", tagValueName)
		}

		switch kind {
		case reflect.String:
			value_.SetString(tagValue)
		case reflect.Bool:
			switch tagValue {
			case "true", "True", "TRUE":
				value_.SetBool(true)
			case "false", "False", "FALSE":
				value_.SetBool(false)
			default:
				return fmt.Errorf("%s is not bool", tagValueName)
			}
		case reflect.Int, reflect.Int64:
			v, err := strconv.ParseInt(tagValue, 10, 64)
			if err != nil {
				return fmt.Errorf("%s is not int", tagValueName)
			}
			value_.SetInt(v)
		case reflect.Float64:
			v, err := strconv.ParseFloat(tagValue, 64)
			if err != nil {
				return fmt.Errorf("%s is not float", tagValueName)
			}
			value_.SetFloat(v)
		}
	}

	return nil
}

func Decode(ptr interface{}) error {
	return decodeWithTagFromDataSource(ptr, "env", &Env{})
}

type Env struct {
}

func (Env) Get(key string) string {
	return os.Getenv(key)
}
