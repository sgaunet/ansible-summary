package ansiblesummary

type Stat struct {
	Changed     int `json:"changed"`
	Failures    int `json:"failures"`
	Ignored     int `json:"ignored"`
	Ok          int `json:"ok"`
	Rescued     int `json:"rescued"`
	Skipped     int `json:"skipped"`
	Unreachable int `json:"unreachable"`
}

type Play struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

type Task struct {
	Name string `json:"name"`
}

type Host struct {
	Changed bool `json:"changed"`
}

type Tasks struct {
	Hosts map[string]Host `json:"hosts"`
	Task  Task            `json:"task"`
}

type Plays struct {
	Play  Play    `json:"play"`
	Tasks []Tasks `json:"tasks"`
}

type AnsibleSummary struct {
	Plays []Plays         `json:"plays"`
	Stats map[string]Stat `json:"stats"`
}
