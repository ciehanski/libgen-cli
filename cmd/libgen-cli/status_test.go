package libgen_cli

import (
	"bytes"
	"io/ioutil"
	"strings"
	"testing"
)

func TestStatus(t *testing.T) {
	t.Skip()
	// Create command
	cmd := statusCmd
	b := bytes.NewBufferString("")
	// Set command output to our bytes
	cmd.SetOut(b)
	if err := cmd.Execute(); err != nil {
		t.Fatal(err)
	}
	// Read bytes outputted by command
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	// Confirm values are expected
	if strings.Contains(string(out), "gen.lib.rus.ec") {
		t.Fatalf("expected \"%s\" got \"%s\"", "gen.lib.rus.ec", string(out))
	}
}
