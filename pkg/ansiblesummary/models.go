package ansiblesummary

// AnsibleSummary represents the struct of the json output
// of the ansible-playbook command if env var below are positionned
// export ANSIBLE_CALLBACKS_ENABLED=json
// export ANSIBLE_STDOUT_CALLBACK=json
type AnsibleSummary struct {
	Plays []Plays         `json:"plays"`
	Stats map[string]Stat `json:"stats"` // one stat by hostname
}

// Plays is the struct to describe a whole run of playbook
// There is basic playbook information (Play)
// And the tasks
type Plays struct {
	Play  Play    `json:"play"`
	Tasks []Tasks `json:"tasks"`
}

// Play represents informations about playbook (omit duration, not needed)
type Play struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

// Tasks represents the tasks of playbook
type Tasks struct {
	Hosts map[string]Host `json:"hosts"`
	Task  Task            `json:"task"`
}

// Task represents informations about task of playbook (omit informations)
type Task struct {
	Name string `json:"name"`
}

// Stats of tasks of the playbook for one host
type Stat struct {
	Changed     int `json:"changed"`
	Failures    int `json:"failures"`
	Ignored     int `json:"ignored"`
	Ok          int `json:"ok"`
	Rescued     int `json:"rescued"`
	Skipped     int `json:"skipped"`
	Unreachable int `json:"unreachable"`
}

type Host struct {
	Changed bool `json:"changed"`
}
