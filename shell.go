package shell

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"

	"github.com/kballard/go-shellquote"
)

// Runf a command in the shell. See Run for more details.
func Runf(command string, a ...interface{}) string {
	return Run(fmt.Sprintf(command, a...))
}

// Run a command in the shell. This command will exit immediately with status
// code 1 if it encounters any errors. The commands run in the current working
// directory with the environment passed in. Standard error goes directly out to
// the shell. Stdout is buffered returned when the command finishes.
func Run(command string) (stdout string) {
	words, err := shellquote.Split(command)
	if err != nil {
		fatal(fmt.Sprintf(`unable to split command: %v`, err))
	}
	if len(words) == 0 {
		fatal(`provided command is empty`)
	}
	result, err := run(words[0], words[1:]...)
	if err != nil {
		fatal(fmt.Sprintf(`error running command %q: %v`, command, err))
	}
	return result
}

// Run the command in the given environment
func run(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	cmd.Env = os.Environ()
	cmd.Stderr = os.Stderr
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return stdout.String(), nil
}

// Print out a message
func fatal(message string) {
	fmt.Fprintf(os.Stderr, "shell: %s\n", message)
	os.Exit(1)
}
