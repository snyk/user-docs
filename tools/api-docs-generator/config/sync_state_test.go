package config

import (
	"reflect"
	"testing"
)

func TestParseSyncState(t *testing.T) {
	tests := []struct {
		name    string
		args    func() string
		want    SyncStateConfig
		wantErr bool
	}{
		{
			name: "parses the config file",
			args: func() string {
				return createTempFile(t, `lastSyncedVersion: 2024-04-25`)
			},
			want: SyncStateConfig{LastSyncedVersion: "2024-04-25"},
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
			got, err := LoadSyncState(tt.args())
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
