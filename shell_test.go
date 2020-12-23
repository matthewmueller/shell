package shell_test

import (
	"testing"

	"github.com/matthewmueller/shell"
	"github.com/tj/assert"
)

var sh = shell.Runf

func TestNoArgs(t *testing.T) {
	stdout := sh("echo")
	assert.Equal(t, "\n", stdout)
}

func TestEcho(t *testing.T) {
	hi := sh("echo %q", "hi")
	assert.Equal(t, "hi\n", hi)
}

func TestEnv(t *testing.T) {
	stdout := sh(`echo -n $GOPATH`)
	assert.NotEmpty(t, stdout)
}
