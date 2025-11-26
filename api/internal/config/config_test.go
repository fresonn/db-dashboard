package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidatePersistentConfig(t *testing.T) {

	t.Parallel()

	tests := []struct {
		name    string
		input   PersistentConfigV1
		wantErr bool
	}{
		{
			name: "valid config",
			input: PersistentConfigV1{
				Connection: Connection{
					Host:     "localhost",
					Port:     5432,
					User:     "admin",
					Password: "pass",
					Database: "mydb",
				},
			},
			wantErr: false,
		},
		{
			name: "missing host",
			input: PersistentConfigV1{
				Connection: Connection{Port: 5432, User: "u", Password: "p", Database: "d"},
			},
			wantErr: true,
		},
		{
			name: "invalid port",
			input: PersistentConfigV1{
				Connection: Connection{Host: "h", Port: 0, User: "u", Password: "p", Database: "d"},
			},
			wantErr: true,
		},
		{
			name: "missing user",
			input: PersistentConfigV1{
				Connection: Connection{Host: "h", Port: 5432, Password: "pass", Database: "d"},
			},
			wantErr: true,
		},
		{
			name: "missing password",
			input: PersistentConfigV1{
				Connection: Connection{Host: "localhost", Port: 5432, User: "admin", Database: "mydb"},
			},
			wantErr: true,
		},
		{
			name: "missing database",
			input: PersistentConfigV1{
				Connection: Connection{Host: "localhost", Port: 5432, User: "admin", Password: "pass"},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErr := ValidatePersistentConfig(&tt.input)
			assert.Equal(t, tt.wantErr, gotErr != nil, "ValidatePersistentConfig() err = %v, wantErr %v", gotErr, tt.wantErr)
		})
	}
}

func TestConfigFromEnvs(t *testing.T) {

	t.Parallel()

	type args struct {
		envs EnvironmentVars
	}

	tests := []struct {
		name        string
		args        args
		wantCfg     bool
		wantErr     bool
		description string
	}{
		{
			name:        "empty envs",
			args:        args{envs: EnvironmentVars{}},
			wantCfg:     false,
			wantErr:     false,
			description: "Should return nil, nil when no envs are set",
		},
		{
			name: "partial envs",
			args: args{
				envs: EnvironmentVars{PgHost: "localhost", PgUser: "admin"},
			},
			wantCfg:     false,
			wantErr:     true,
			description: "Should return error if some but not all envs are set",
		},
		{
			name: "all envs set",
			args: args{
				envs: EnvironmentVars{
					PgHost:     "localhost",
					PgPort:     5432,
					PgUser:     "admin",
					PgPassword: "secret",
					PgDatabase: "db",
				},
			},
			wantCfg:     true,
			wantErr:     false,
			description: "Returns config because all PG_ envs were set",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := configFromEnvs(tt.args.envs)

			assert.Equal(t, tt.wantErr, err != nil, "error: %v (wantErr=%v)", err, tt.wantErr)
			assert.Equal(t, tt.wantCfg, got != nil, "config: %+v (wantCfg=%v)", got, tt.wantCfg)
		})
	}
}

func TestMigrate(t *testing.T) {

	t.Parallel()

	tests := []struct {
		name        string
		inputCfg    *PersistentConfigV1
		wantChanged bool
		wantVersion int
	}{
		{
			name: "old version",
			inputCfg: &PersistentConfigV1{
				Version: 0,
			},
			wantChanged: true,
			wantVersion: 1,
		},
		{
			name: "current version",
			inputCfg: &PersistentConfigV1{
				Version: 1,
			},
			wantChanged: false,
			wantVersion: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, changed := Migrate(tt.inputCfg)

			assert.Equal(t, tt.wantChanged, changed, "changed flag does not match")
			assert.Equal(t, tt.inputCfg, got, "the migration result does not match")
		})
	}
}

func cleanUp() {
	_ = os.RemoveAll("./data")
}

func TestLoadPersistentConfig(t *testing.T) {
	// We do not use t.Parallel(), since we are working with one folder ./data
	//
	t.Cleanup(cleanUp)

	tests := []struct {
		name    string
		input   *PersistentConfigV1
		wantErr bool
	}{
		{
			name: "standard config",
			input: &PersistentConfigV1{
				Version: 1,
				Connection: Connection{
					Host:     "127.0.0.1",
					Port:     5432,
					User:     "admin",
					Password: "secure_password",
					Database: "production_db",
					SSLMode:  "disable",
				},
			},
			wantErr: false,
		},
		{
			name: "not found or broken config",
			input: &PersistentConfigV1{
				Version: 1,
				Connection: Connection{
					Host:    "127.0.0.1",
					Port:    5432,
					SSLMode: "disable",
				},
			},
			wantErr: false,
		},
		{
			name: "config with special characters",
			input: &PersistentConfigV1{
				Version: 1,
				Connection: Connection{
					Host:     "db.example.com",
					Port:     1234,
					User:     "user with spaces",
					Password: "p@ssw0rd!\"#$%&'()*+,",
					Database: "weird-db-name",
					SSLMode:  "require",
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cleanUp() // before each test

			err := tt.input.Save()
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err, "save failed")

			loadedCfg, err := LoadPersistentConfig()
			require.NoError(t, err, "load failed")

			assert.Equal(t, tt.input, loadedCfg, "loaded config doesn't match saved config")
		})
	}
}

func TestNew_WithBrokenConfigFile(t *testing.T) {
	t.Cleanup(cleanUp)
	cleanUp()

	// Создаем битый JSON
	dir := filepath.Dir(ConfigPath)
	err := os.MkdirAll(dir, 0o755)
	require.NoError(t, err)

	err = os.WriteFile(ConfigPath, []byte("{ broken json "), 0o600)
	require.NoError(t, err)

	// Очищаем ENV, чтобы New() не подхватил ничего лишнего
	os.Clearenv()

	// Запускаем тестируемую функцию
	appCfg := New()

	// Проверяем результаты
	assert.Nil(t, appCfg.PersistentConfig, "PersistentConfig should be nil when file is broken")

	// Проверяем, что файл был удален
	_, err = os.Stat(ConfigPath)
	assert.True(t, os.IsNotExist(err), "Broken config file should have been deleted")
}
