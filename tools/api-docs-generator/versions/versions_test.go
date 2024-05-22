package versions

import "testing"

func Test_getLatestGAVersion(t *testing.T) {
	tests := []struct {
		name     string
		versions []string
		want     string
	}{
		{
			name: "gets the latest version",
			versions: []string{
				"2024-04-29~experimental", "2024-04-29~beta", "2024-04-29",
				"2024-05-08~experimental", "2024-05-08~beta", "2024-05-08"},
			want: "2024-05-08",
		},
		{
			name: "gets the latest GA version",
			versions: []string{
				"2024-04-25~experimental", "2024-04-25~beta", "2024-04-25",
				"2024-04-29~experimental", "2024-04-29~beta", "2024-04-29",
				"2024-05-08~experimental", "2024-05-08~beta"},
			want: "2024-04-29",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetLatestGAVersion(tt.versions); got != tt.want {
				t.Errorf("getLatestGAVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}
