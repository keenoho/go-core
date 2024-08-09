package core_test

import (
	"os"
	"testing"

	"github.com/keenoho/go-core"
)

func TestConfig(t *testing.T) {
	core.ConfigLoad()
	t.Log(os.Environ())
	core.ConfigSet("foo", "bar")
	t.Log(os.Getenv("foo"))
	core.ConfigSetMap(map[string]string{"foo2": "bar2"})
	t.Log(os.Getenv("foo2"))
}

func TestDevCofig(t *testing.T) {
	core.ConfigLoad(core.ConfigOption{Env: "development"})
	t.Log(os.Environ())
}

func TestDirCofig(t *testing.T) {
	pwd, _ := os.Getwd()
	core.ConfigLoad(core.ConfigOption{Env: "development", EnvDir: pwd})
	t.Log(os.Environ())
}
