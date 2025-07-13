package ansiblesummary

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// Output interface defines methods for writing Ansible summary statistics.
type Output interface {
	WriteStats(a *AnsibleSummary) []error
	WriteStatsHTML(a *AnsibleSummary) []error
	WriteStatsJSON(a *AnsibleSummary) error
}

type output struct {
	w io.Writer
}

// NewOutput creates a new Output instance that writes to stdout.
// Returns concrete type to allow method calls like SetOutput.
//
//nolint:revive // Need concrete type for SetOutput method
func NewOutput() *output {
	return &output{
		w: os.Stdout,
	}
}

func (o *output) SetOutput(w io.Writer) {
	o.w = w
}

func (o *output) WriteStats(a *AnsibleSummary) []error {
	var errs []error
	for hostname, status := range a.Stats {
		_, err := fmt.Fprintf(o.w,
			"%-50s  ok=%2d  changed=%d  unreachable=%2d failures=%2d skipped=%2d rescued=%2d ignored=%2d\n",
			hostname, status.Ok, status.Changed, status.Unreachable, status.Failures,
			status.Skipped, status.Rescued, status.Ignored)
		if err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}

func (o *output) WriteStatsHTML(a *AnsibleSummary) []error {
	var errs []error
	for hostname, status := range a.Stats {
		_, err := fmt.Fprintf(o.w, "%-50s  %s %s %s %s %s %s %s\n",
			hostname,
			addColor("ok", status.Ok, "green"),
			addColor("changed", status.Changed, "orange"),
			addColor("unreachable", status.Unreachable, "red"),
			addColor("failures", status.Failures, "red"),
			addColor("skipped", status.Skipped, "blue"),
			addColor("rescued", status.Rescued, "orange"),
			addColor("ignored", status.Ignored, "black"))
		if err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}

func (o *output) WriteStatsJSON(a *AnsibleSummary) error {
	jsonData, err := json.MarshalIndent(a.Stats, "", "    ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}
	_, err = fmt.Fprintf(o.w, "%s\n", string(jsonData))
	if err != nil {
		return fmt.Errorf("failed to write output: %w", err)
	}
	return nil
}
