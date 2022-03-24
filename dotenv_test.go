package dotenv

import "testing"

func TestDotEnv(t *testing.T) {
	Load()
	if Get("TEST_VAR") != "zero" {
		t.Error("Expected TEST_VAR to be `test`")
	}
}
