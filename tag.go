package dotenv

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
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

		attribute := newAttribute(typ.Tag, tagName, typ.Name)
		tagValue, err := attribute.Value(dataSource.Get(attribute.Name))
		if err != nil {
			return err
		}

		if err := setValue(typ.Type, val, tagValue, attribute.Name); err != nil {
			return err
		}
	}

	return nil
}

type attribute struct {
	Name     string
	Required bool
	Default  string
}

func (a *attribute) Value(current string) (string, error) {
	if current == "" {
		if a.Default != "" {
			return a.Default, nil
		}

		if a.Required {
			return "", fmt.Errorf("%s is required", a.Name)
		}
	}

	return current, nil
}

func newAttribute(tag reflect.StructTag, tagName string, attributeName string) *attribute {
	tagValueName := tag.Get(tagName)
	tabValueDfeault := tag.Get("default")
	tagValueRequired := tag.Get("required")

	if tagValueName == "" {
		tagValueName = strings.ToUpper(attributeName)
	}

	return &attribute{
		Name:     tagValueName,
		Required: tagValueRequired == "true",
		Default:  tabValueDfeault,
	}
}

func setValue(t reflect.Type, v reflect.Value, Value string, Name string) error {
	switch v.Kind() {
	case reflect.String:
		v.SetString(Value)

	case reflect.Bool:
		switch Value {
		case "true", "True", "TRUE":
			v.SetBool(true)
		case "false", "False", "FALSE":
			v.SetBool(false)
		default:
			return fmt.Errorf("%s is not bool", Name)
		}

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		vv, err := strconv.ParseInt(Value, 10, 64)
		if err != nil {
			return fmt.Errorf("%s is not int64", Name)
		}
		v.SetInt(vv)

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		vv, err := strconv.ParseUint(Value, 10, 64)
		if err != nil {
			return fmt.Errorf("%s is not int64", Name)
		}
		v.SetUint(vv)

	case reflect.Float32, reflect.Float64:
		vv, err := strconv.ParseFloat(Value, 64)
		if err != nil {
			return fmt.Errorf("%s is not float", Name)
		}
		v.SetFloat(vv)

	default:
		return fmt.Errorf("%s is not supported", Name)
	}

	return nil
}

// Decode decodes the given struct pointer from the environment.
func Decode(ptr interface{}) error {
	return decodeWithTagFromDataSource(ptr, "env", &EnvDataSource{})
}

// EnvDataSource is a data source that loads data from the environment.
type EnvDataSource struct {
}

// Get returns the value of the given key.
func (EnvDataSource) Get(key string) string {
	return os.Getenv(key)
}
