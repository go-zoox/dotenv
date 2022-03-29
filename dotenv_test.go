package dotenv

import (
	"os"
	"testing"
)

func TestDotEnv(t *testing.T) {
	// Load()
	// if Get("TEST_VAR") != "zero" {
	// 	t.Error("Expected TEST_VAR to be `test`")
	// }

	type Env struct {
		User    string `env:"USER" required:"true"`
		HomeDir string `env:"HOME" required:"true"`
		Age     int    `env:"AGE" default:"100"`
	}

	var env Env
	if err := Load(&env); err != nil {
		t.Errorf("Expected Load() to not return an error")
	}

	if env.User != os.Getenv("USER") {
		t.Errorf("Expected %s, got %s", os.Getenv("USER"), env.User)
	}

	if env.HomeDir != os.Getenv("HOME") {
		t.Errorf("Expected %s, got %s", os.Getenv("HOME"), env.HomeDir)
	}

	if env.Age != 100 {
		t.Errorf("Expected %d, got %d", 100, env.Age)
	}

	// Get()
	if Get("USER") != os.Getenv("USER") {
		t.Errorf("Expected %s, got %s", os.Getenv("USER"), Get("USER"))
	}

	if Get("HOME") != os.Getenv("HOME") {
		t.Errorf("Expected %s, got %s", os.Getenv("HOME"), Get("HOME"))
	}

	if Get("PWD") != os.Getenv("PWD") {
		t.Errorf("Expected %s, got %s", os.Getenv("PWD"), Get("PWD"))
	}

	if Get("TEST_VAR") != "zero" {
		t.Errorf("Expected %s, got %s", "zero", Get("TEST_VAR"))
	}

	if Get("NONE_EXIST") != "" {
		t.Errorf("Expected %s, got %s", "", Get("NONE_EXIST"))
	}
}
