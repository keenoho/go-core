package core_test

import (
	"testing"

	"github.com/keenoho/go-core"
)

func TestConfig(t *testing.T) {
	core.ConfigLoad()
	env := core.ConfigGet("ENV")
	t.Log(env)
}
