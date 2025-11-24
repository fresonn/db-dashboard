package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_loadFromFile(t *testing.T) {
	tests := []struct {
		name    string
		content string
		want    *PersistentConfigV1
		wantErr bool
	}{
		{
			name: "valid config",
			content: `{
                "version": 1,
                "connection": {
                    "host": "prod-db.example.com",
                    "port": 5432,
                    "user": "app",
                    "password": "supersecret",
                    "database": "app_production",
                    "sslmode": "require"
                }
            }`,
			want: &PersistentConfigV1{
				Version: 1,
				Connection: Connection{
					Host:     "prod-db.example.com",
					Port:     5432,
					User:     "app",
					Password: "supersecret",
					Database: "app_production",
					SSLMode:  "require",
				},
			},
			wantErr: false,
		},
		{
			name:    "file does not exist",
			content: "", // no file
			want:    nil,
			wantErr: true,
		},
		{
			name:    "invalid json",
			content: `{"version": 1, "connection": {invalid json!}`,
			want:    nil,
			wantErr: true,
		},
		{
			name:    "empty file",
			content: ``,
			want:    nil,
			wantErr: true,
		},
		{
			name:    "trailing comma / extra data",
			content: `{"version": 1, "connection": {"host":"localhost","port":5432,}}`,
			want:    nil,
			wantErr: true,
		},
		{
			name: "minimal valid config (only required fields)",
			content: `{
                "version": 1,
                "connection": {
                    "host": "localhost",
                    "port": 5432,
                    "user": "",
                    "password": "",
                    "database": "",
                    "sslmode": "disable"
                }
            }`,
			want: &PersistentConfigV1{
				Version: 1,
				Connection: Connection{
					Host:    "localhost",
					Port:    5432,
					SSLMode: "disable",
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var path string

			if tt.content != "" {
				// create a temporary file only if there is content
				file, err := os.CreateTemp("", "config-*.json")
				assert.NoError(t, err, "failed to create temporary file")

				defer os.Remove(file.Name())

				_, err = file.WriteString(tt.content)
				assert.NoError(t, err)
				assert.NoError(t, file.Close())

				path = file.Name()
			} else {
				// for the case "file does not exist" - we take a non-existent path
				path = filepath.Join(t.TempDir(), "non-existent-config.json")
			}

			got, err := loadFromFile(path)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, got)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
