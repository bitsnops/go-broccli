package broccli

import (
	"os"
	"testing"
	"fmt"
)

func TestCLI(t *testing.T) {
	c := NewCLI("Example", "App", "Author <a@example.com>")
	cmd1 := c.AddCmd("cmd1", "Prints out a string", func(c *CLI) int {
		fmt.Fprintf(os.Stdout, "Printed out %s%s\n\n", c.Flag("tekst"), c.Flag("alphanumdots"))
		return 2
	})
	cmd1.AddFlag("tekst", "t", "Text", "Text to print", TypeString, IsRequired)
	cmd1.AddFlag("alphanumdots", "a", "Alphanum with dots", "Can have dots", TypeAlphanumeric, AllowDots)

	os.Args = []string{"test", "cmd1"}
	got := c.Run()
	if got != 1 {
		t.Errorf("CLI.Run() should have returned 1 instead of %d", got)
	}

	os.Args = []string{"test", "cmd1", "-t", ""}
	got = c.Run()
	if got != 1 {
		t.Errorf("CLI.Run() should have returned 1 instead of %d", got)
	}

	os.Args = []string{"test", "cmd1", "--tekst", "Tekst123", "--alphanumdots"}
	got = c.Run()
	if got != 2 {
		t.Errorf("CLI.Run() should have returned 2 instead of %d", got)
	}

	os.Args = []string{"test", "cmd1", "--tekst", "Tekst123", "--alphanumdots", "aZ0-9"}
	got = c.Run()
	if got != 1 {
		t.Errorf("CLI.Run() should have returned 1 instead of %d", got)
	}

	os.Args = []string{"test", "cmd1", "--tekst", "Tekst123", "--alphanumdots", "aZ0.9"}
	got = c.Run()
	if got != 2 {
		t.Errorf("CLI.Run() should have returned 2 instead of %d", got)
	}

	// check output
	// check post validation for both failure and success
	// check onTrue
}
