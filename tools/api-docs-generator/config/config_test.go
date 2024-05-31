package config

import (
	"os"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name    string
		args    func() string
		want    *Config
		wantErr bool
	}{
		{
			name: "parses the config file",
			args: func() string {
				return createTempFile(t, `fetcher:
  source: source
  destination: destination
specs:
- path: .gitbook/assets/spec.yaml
  suffix: " (v1)"
  docsHint: hint 1
- path: .gitbook/assets/rest-spec.json
  docsHint: hint 2

output:
  apiReferencePath: snyk-api/reference`)
			},
			want: &Config{
				Fetcher: Fetcher{"source", "destination"},
				Specs: []Spec{
					{".gitbook/assets/spec.yaml", " (v1)", "hint 1"},
					{".gitbook/assets/rest-spec.json", "", "hint 2"},
				},
				Output: Output{APIReferencePath: "snyk-api/reference"},
			},
		},
		{
			name: "invalid config",
			args: func() string {
				return createTempFile(t, `test 12`)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args())
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func createTempFile(t *testing.T, content string) string {
	t.Helper()
	file, err := os.CreateTemp("", "config")
	if err != nil {
		t.Fatal(err)
	}
	_, err = file.WriteString(content)
	if err != nil {
		t.Fatal(err)
	}
	return file.Name()
}
