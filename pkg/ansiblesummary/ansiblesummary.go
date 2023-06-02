package ansiblesummary

import (
	"encoding/json"
	"fmt"
	"os"
)

func NewAnsibleSummaryFromFile(filename string) (*AnsibleSummary, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var r AnsibleSummary
	err = json.NewDecoder(f).Decode(&r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

func (a *AnsibleSummary) PrintStats() {
	for hostname, status := range a.Stats {
		fmt.Printf("%-50s  ok=%2d  changed=%d  unreachable=%2d failures=%2d skipped=%2d rescued=%2d ignored=%2d\n",
			hostname, status.Ok, status.Changed, status.Unreachable, status.Failures, status.Skipped, status.Rescued, status.Ignored)
	}
}

func (a *AnsibleSummary) PrintHTMLStats() {
	for hostname, status := range a.Stats {
		fmt.Printf("%-50s  %s %s %s %s %s %s %s\n",
			hostname, addColor("ok", status.Ok, "green"), addColor("changed", status.Changed, "orange"), addColor("unreachable", status.Unreachable, "red"),
			addColor("failures", status.Failures, "red"), addColor("skipped", status.Skipped, "blue"),
			addColor("rescued", status.Rescued, "orange"), addColor("ignored", status.Ignored, "black"))
	}
}

// function addColor return a string colorized only if value > 0
func addColor(prefixstr string, value int, color string) string {
	if value > 0 {
		return fmt.Sprintf("<span style=\"color:%s\">%s=%d</span>", color, prefixstr, value)
	} else {
		return fmt.Sprintf("%s=%d", prefixstr, value)
	}
}

func (a *AnsibleSummary) PrintNameOfTaksNotOK() {
	var i int
	for _, p := range a.Plays {
		for _, t := range p.Tasks {
			for h := range t.Hosts {
				if t.Hosts[h].Changed {
					i++
					if i == 1 {
						fmt.Println("Tasks not synchronised :")
					}
					fmt.Printf("On Host %30s task %s is not synchronised.\n", h, t.Task.Name)
				}
			}
		}
	}
}
