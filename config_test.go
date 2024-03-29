package core_test

import (
	"os"
	"testing"

	"github.com/keenoho/go-core"
)

func TestConfig(t *testing.T) {
	core.ConfigLoad("development")
	env := core.ConfigGet("ENV")
	host := core.ConfigGet("HOST")
	port := core.ConfigGet("PORT")
	t.Log(env, host, port)
	t.Log(os.Environ())
}
