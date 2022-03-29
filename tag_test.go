package dotenv

import (
	"os"
	"testing"
)

func TestTag(t *testing.T) {
	type Cat struct {
		Name string `env:"USER"`
		Age  int    `env:"AGE" default:"100"`
		Type string `env:"TYPE"`
	}

	var cat Cat

	Decode(&cat)

	if cat.Name != os.Getenv("USER") {
		t.Errorf("Expected %s, got %s", os.Getenv("USER"), cat.Name)
	}

	if cat.Age != 100 {
		t.Errorf("Expected %d, got %d", 100, cat.Age)
	}

	if cat.Type != os.Getenv("TYPE") {
		t.Errorf("Expected %s, got %s", "type", cat.Type)
	}

	// fmt.Println("xxxx:", cat)
}
