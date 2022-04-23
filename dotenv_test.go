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
		User    string `env:"USER,required"`
		HomeDir string `env:"HOME,required"`
		Age     int    `env:"AGE,default=100"`

		// Database struct {
		// 	Host string `env:"DB_HOST" default:"127.0.0.1"`
		// 	Port int    `env:"DB_PORT" default:"3306"`
		// 	User string `env:"DB_USER" default:"zero"`
		// 	Pass string `env:"DB_PASS" default:"zero"`
		// 	Name string `env:"DB_NAME" default:"zero"`
		// }
	}

	var env Env
	if err := Load(&env); err != nil {
		t.Errorf("Expected Load() to not return an error: %s", err.Error())
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

	// // Nest
	// if env.Database.Host != "127.0.0.1" {
	// 	t.Errorf("Expected %s, got %s", "127.0.0.1", env.Database.Host)
	// }

	// if env.Database.Port != 3306 {
	// 	t.Errorf("Expected %d, got %d", 3306, env.Database.Port)
	// }

	// if env.Database.User != "zero" {
	// 	t.Errorf("Expected %s, got %s", "zero", env.Database.User)
	// }

	// if env.Database.Pass != "zero" {
	// 	t.Errorf("Expected %s, got %s", "zero", env.Database.Pass)
	// }

	// if env.Database.Name != "zero" {
	// 	t.Errorf("Expected %s, got %s", "zero", env.Database.Name)
	// }

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
