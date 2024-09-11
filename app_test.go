package core_test

import (
	"testing"

	"github.com/keenoho/go-core"
)

func TestHttpServer(t *testing.T) {
	core.ConfigLoad()
	app := core.AppNew()
	app.RegisterController(
		new(TestHttpController),
	)
	err := app.Start()
	if err != nil {
		t.Fatal(err)
	}
}
