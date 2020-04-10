package libgen_cli

import (
	"bytes"
	"io/ioutil"
	"strings"
	"testing"
)

func TestLink(t *testing.T) {
	// Create command
	cmd := linkCmd
	b := bytes.NewBufferString("")
	// Set command output to our bytes
	cmd.SetOut(b)
	// Add arguments
	cmd.SetArgs([]string{"2F2DBA2A621B693BB95601C16ED680F8"})
	if err := cmd.Execute(); err != nil {
		t.Fatalf("error executing command: %v", err)
	}
	// Read bytes outputted by command
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatalf("error reading command output: %v", err)
	}
	// Confirm values are expected
	if strings.Contains(string(out), "https://b-ok.cc/dl/") {
		t.Fatalf("expected \"%s\" got \"%s\"", "https://b-ok.cc/dl/", string(out))
	}
}
