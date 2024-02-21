package ansiblesummary_test

import (
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
