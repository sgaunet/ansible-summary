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

// function addColor return a string colorized only if value > 0
func addColor(prefixstr string, value int, color string) string {
	if value > 0 {
		return fmt.Sprintf("<span style=\"color:%s\">%s=%d</span>", color, prefixstr, value)
	} else {
		return fmt.Sprintf("%s=%d", prefixstr, value)
	}
}

func (a *AnsibleSummary) GetListOfTasksChanged() []string {
	return []string{"to implement"}
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
