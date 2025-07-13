package ansiblesummary_test

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/sgaunet/ansible-summary/pkg/ansiblesummary"
	"github.com/stretchr/testify/assert"
)

func TestNewAnsibleSummaryFromFile(t *testing.T) {
	type args struct {
		content string
	}
	tests := []struct {
		name    string
		args    args
		want    *ansiblesummary.AnsibleSummary
		wantErr bool
	}{
		{
			name: "empty file",
			args: args{
				content: "",
			},
			want:    &ansiblesummary.AnsibleSummary{},
			wantErr: true,
		},
		{
			name: "empty json",
			args: args{
				content: "{}",
			},
			want:    &ansiblesummary.AnsibleSummary{},
			wantErr: false,
		},
		{
			name: "normal case",
			args: args{
				content: `
			{
				"stats": {
					"alma9.2": {
						"changed": 0,
						"failures": 0,
						"ignored": 0,
						"ok": 4,
						"rescued": 0,
						"skipped": 1,
						"unreachable": 0
					},
					"rocky9.2": {
						"changed": 0,
						"failures": 0,
						"ignored": 0,
						"ok": 4,
						"rescued": 0,
						"skipped": 1,
						"unreachable": 0
					}
				}
			}
`,
			},
			want: &ansiblesummary.AnsibleSummary{
				Stats: map[string]ansiblesummary.Stat{
					"alma9.2": {
						Changed:     0,
						Failures:    0,
						Ignored:     0,
						Ok:          4,
						Rescued:     0,
						Skipped:     1,
						Unreachable: 0,
					},
					"rocky9.2": {
						Changed:     0,
						Failures:    0,
						Ignored:     0,
						Ok:          4,
						Rescued:     0,
						Skipped:     1,
						Unreachable: 0,
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// create temporary file and write content in it
			f, err := os.CreateTemp("/tmp", "test-ansible-summary")
			assert.Nil(t, err)
			defer os.Remove(f.Name())
			_, err = f.Write([]byte(tt.args.content))
			assert.Nil(t, err)
			err = f.Close()
			assert.Nil(t, err)
			got, err := ansiblesummary.NewAnsibleSummaryFromFile(f.Name())
			assert.True(t, (err != nil) == tt.wantErr)
			if !tt.wantErr {
				assert.Equal(t, tt.want, got)
			}
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("NewAnsibleSummaryFromFile() = %v, want %v", got, tt.want)
			// }
		})
	}
}

func TestAnsibleSummary_HasChangedOrFailed(t *testing.T) {
	tests := []struct {
		name string
		data *ansiblesummary.AnsibleSummary
		want bool
	}{
		{
			name: "no changes",
			data: &ansiblesummary.AnsibleSummary{
				Plays: []ansiblesummary.Plays{
					{
						Tasks: []ansiblesummary.Tasks{
							{
								Hosts: map[string]ansiblesummary.Host{
									"host1": {Changed: false},
									"host2": {Changed: false},
								},
							},
						},
					},
				},
			},
			want: false,
		},
		{
			name: "has changes",
			data: &ansiblesummary.AnsibleSummary{
				Plays: []ansiblesummary.Plays{
					{
						Tasks: []ansiblesummary.Tasks{
							{
								Hosts: map[string]ansiblesummary.Host{
									"host1": {Changed: false},
									"host2": {Changed: true},
								},
							},
						},
					},
				},
			},
			want: true,
		},
		{
			name: "empty data",
			data: &ansiblesummary.AnsibleSummary{},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.data.HasChangedOrFailed()
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestAnsibleSummary_GetListOfTasksChanged(t *testing.T) {
	data := &ansiblesummary.AnsibleSummary{}
	result := data.GetListOfTasksChanged()
	expected := []string{"to implement"}
	assert.Equal(t, expected, result)
}

func TestNewOutput(t *testing.T) {
	output := ansiblesummary.NewOutput()
	assert.NotNil(t, output)
}

func TestOutput_WriteStatsJSON(t *testing.T) {
	data := &ansiblesummary.AnsibleSummary{
		Stats: map[string]ansiblesummary.Stat{
			"host1": {
				Ok:          5,
				Changed:     1,
				Failures:    0,
				Unreachable: 0,
				Skipped:     2,
				Rescued:     0,
				Ignored:     0,
			},
		},
	}

	var buf bytes.Buffer
	output := ansiblesummary.NewOutput()
	
	// Use type assertion with the exact method signature from output.go
	if setter, ok := output.(interface{ SetOutput(w io.Writer) }); ok {
		setter.SetOutput(&buf)
		
		err := output.WriteStatsJSON(data)
		assert.NoError(t, err)
		
		result := buf.String()
		assert.Contains(t, result, "host1")
		assert.Contains(t, result, `"ok": 5`)
		assert.Contains(t, result, `"changed": 1`)
	} else {
		// Test that the method doesn't error even if we can't capture output
		err := output.WriteStatsJSON(data)
		assert.NoError(t, err)
	}
}

func TestOutput_WriteStatsHTML(t *testing.T) {
	data := &ansiblesummary.AnsibleSummary{
		Stats: map[string]ansiblesummary.Stat{
			"host1": {
				Ok:          5,
				Changed:     1,
				Failures:    0,
				Unreachable: 0,
				Skipped:     2,
				Rescued:     0,
				Ignored:     0,
			},
		},
	}

	var buf bytes.Buffer
	output := ansiblesummary.NewOutput()
	
	// Use type assertion with the exact method signature
	if setter, ok := output.(interface{ SetOutput(w io.Writer) }); ok {
		setter.SetOutput(&buf)
		
		errs := output.WriteStatsHTML(data)
		assert.Empty(t, errs)
		
		result := buf.String()
		assert.Contains(t, result, "host1")
		assert.Contains(t, result, "ok=5")
		assert.Contains(t, result, "changed=1")
	} else {
		// Test that the method doesn't error even if we can't capture output
		errs := output.WriteStatsHTML(data)
		assert.Empty(t, errs)
	}
}

func TestOutput_WriteStats(t *testing.T) {
	data := &ansiblesummary.AnsibleSummary{
		Stats: map[string]ansiblesummary.Stat{
			"host1": {
				Ok:          5,
				Changed:     1,
				Failures:    0,
				Unreachable: 0,
				Skipped:     2,
				Rescued:     0,
				Ignored:     0,
			},
		},
	}

	var buf bytes.Buffer
	output := ansiblesummary.NewOutput()
	
	// Use type assertion with the exact method signature
	if setter, ok := output.(interface{ SetOutput(w io.Writer) }); ok {
		setter.SetOutput(&buf)
		
		errs := output.WriteStats(data)
		assert.Empty(t, errs)
		
		result := buf.String()
		assert.Contains(t, result, "host1")
		assert.Contains(t, result, "ok= 5")
		assert.Contains(t, result, "changed=1")
	} else {
		// Test that the method doesn't error even if we can't capture output
		errs := output.WriteStats(data)
		assert.Empty(t, errs)
	}
}

func TestAnsibleSummary_PrintNameOfTaksNotOK(t *testing.T) {
	tests := []struct {
		name string
		data *ansiblesummary.AnsibleSummary
	}{
		{
			name: "with changed tasks",
			data: &ansiblesummary.AnsibleSummary{
				Plays: []ansiblesummary.Plays{
					{
						Tasks: []ansiblesummary.Tasks{
							{
								Task: ansiblesummary.Task{Name: "test task"},
								Hosts: map[string]ansiblesummary.Host{
									"host1": {Changed: true},
									"host2": {Changed: false},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "no changed tasks",
			data: &ansiblesummary.AnsibleSummary{
				Plays: []ansiblesummary.Plays{
					{
						Tasks: []ansiblesummary.Tasks{
							{
								Task: ansiblesummary.Task{Name: "test task"},
								Hosts: map[string]ansiblesummary.Host{
									"host1": {Changed: false},
									"host2": {Changed: false},
								},
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// This method prints to stdout, so we just test it doesn't panic
			assert.NotPanics(t, func() {
				tt.data.PrintNameOfTaksNotOK()
			})
		})
	}
}
