package ansiblesummary

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Output interface {
	WriteStats(a *AnsibleSummary) (errs []error)
	WriteStatsHTML(a *AnsibleSummary) (errs []error)
	WriteStatsJSON(a *AnsibleSummary) (err error)
}

type output struct {
	w io.Writer
}

func NewOutput() Output {
	return &output{
		w: os.Stdout,
	}
}

func (o *output) SetOutput(w io.Writer) {
	o.w = w
}

func (o *output) WriteStats(a *AnsibleSummary) (errs []error) {
	for hostname, status := range a.Stats {
		_, err := fmt.Fprintf(o.w, "%-50s  ok=%2d  changed=%d  unreachable=%2d failures=%2d skipped=%2d rescued=%2d ignored=%2d\n",
			hostname, status.Ok, status.Changed, status.Unreachable, status.Failures, status.Skipped, status.Rescued, status.Ignored)
		if err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}

func (o *output) WriteStatsHTML(a *AnsibleSummary) (errs []error) {
	for hostname, status := range a.Stats {
		_, err := fmt.Fprintf(o.w, "%-50s  %s %s %s %s %s %s %s\n",
			hostname, addColor("ok", status.Ok, "green"), addColor("changed", status.Changed, "orange"), addColor("unreachable", status.Unreachable, "red"),
			addColor("failures", status.Failures, "red"), addColor("skipped", status.Skipped, "blue"),
			addColor("rescued", status.Rescued, "orange"), addColor("ignored", status.Ignored, "black"))
		if err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}

func (o *output) WriteStatsJSON(a *AnsibleSummary) (err error) {
	jsonData, err := json.MarshalIndent(a.Stats, "", "    ")
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(o.w, "%s\n", string(jsonData))
	return err
}
