package main

import (
	"bytes"
	"testing"
)

func TestPushSwapRun(t *testing.T) {
	tests := []struct {
		name        string
		arguments   []string
		wantOutput  string
		wantError   string
		wantFailure bool
	}{
		{name: "no arguments stays silent"},
		{name: "already sorted stays silent", arguments: []string{"1 2"}},
		{name: "two reversed values print one instruction", arguments: []string{"2 1"}, wantOutput: "sa\n"},
		{name: "grouped and separate values are accepted", arguments: []string{"3", "1 2"}, wantOutput: "ra\n"},
		{name: "invalid number prints error", arguments: []string{"one"}, wantError: "Error\n", wantFailure: true},
		{name: "duplicate number prints error", arguments: []string{"1 1"}, wantError: "Error\n", wantFailure: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var output, errorOutput bytes.Buffer
			err := run(test.arguments, &output, &errorOutput)

			if (err != nil) != test.wantFailure {
				t.Fatalf("run error = %v, want failure %v", err, test.wantFailure)
			}
			if got := output.String(); got != test.wantOutput {
				t.Errorf("stdout = %q, want %q", got, test.wantOutput)
			}
			if got := errorOutput.String(); got != test.wantError {
				t.Errorf("stderr = %q, want %q", got, test.wantError)
			}
		})
	}
}

func TestPushSwapRunWritesOnlyInstructions(t *testing.T) {
	var output, errorOutput bytes.Buffer
	if err := run([]string{"5 4 3 2 1"}, &output, &errorOutput); err != nil {
		t.Fatalf("run returned an error: %v", err)
	}
	for _, line := range bytes.Split(bytes.TrimSuffix(output.Bytes(), []byte("\n")), []byte("\n")) {
		switch string(line) {
		case "sa", "sb", "ss", "pa", "pb", "ra", "rb", "rr", "rra", "rrb", "rrr":
		default:
			t.Errorf("stdout contains non-instruction line %q", line)
		}
	}
	if errorOutput.Len() != 0 {
		t.Errorf("stderr = %q, want empty output", errorOutput.String())
	}
}
