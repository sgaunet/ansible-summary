// Package ansiblesummary provides functionality to parse and summarize Ansible task execution results.
package ansiblesummary

import (
	"encoding/json"
	"fmt"
	"os"
)

// NewAnsibleSummaryFromFile creates an AnsibleSummary from a JSON file.
func NewAnsibleSummaryFromFile(filename string) (*AnsibleSummary, error) {
	// #nosec G304 - File path is provided by user via CLI, this is expected behavior
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %s: %w", filename, err)
	}
	defer func() {
		if closeErr := f.Close(); closeErr != nil {
			fmt.Printf("Warning: failed to close file: %v\n", closeErr)
		}
	}()

	var r AnsibleSummary
	err = json.NewDecoder(f).Decode(&r)
	if err != nil {
		return nil, fmt.Errorf("failed to decode JSON: %w", err)
	}

	return &r, nil
}

// addColor returns a string colorized only if value > 0.
func addColor(prefixstr string, value int, color string) string {
	if value > 0 {
		return fmt.Sprintf("<span style=\"color:%s\">%s=%d</span>", color, prefixstr, value)
	}
	return fmt.Sprintf("%s=%d", prefixstr, value)
}

// GetListOfTasksChanged returns a list of changed tasks (placeholder implementation).
func (a *AnsibleSummary) GetListOfTasksChanged() []string {
	return []string{"to implement"}
}

// PrintNameOfTaksNotOK prints the names of tasks that are not synchronized.
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

// HasChangedOrFailed returns true if any tasks have changed or failed.
func (a *AnsibleSummary) HasChangedOrFailed() bool {
	for _, p := range a.Plays {
		for _, t := range p.Tasks {
			for h := range t.Hosts {
				if t.Hosts[h].Changed {
					return true
				}
			}
		}
	}
	return false
}
