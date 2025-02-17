package database_test

import (
	"testing"
	"time"

	database "crypto-dashboard/pkg/database/postgres"
	"crypto-dashboard/pkg/settings"
)

func TestNewConnection(t *testing.T) {
	type args struct {
		cfg *settings.SQLSetting
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Valid configuration",
			args: args{
				cfg: &settings.SQLSetting{
					Host:            "localhost",
					Port:            5432,
					Username:        "postgres",
					Password:        "password",
					DBname:          "testdb",
					SSLMode:         "disable",
					MaxConns:        10,
					MaxConnIdleTime: 5,
					MaxConnLifetime: uint64(30 * time.Minute),
				},
			},
			wantErr: false,
		},
		{
			name: "Invalid host",
			args: args{
				cfg: &settings.SQLSetting{
					Host:            "invalid_host",
					Port:            5432,
					Username:        "postgres",
					Password:        "password",
					DBname:          "testdb",
					SSLMode:         "disable",
					MaxConns:        10,
					MaxConnIdleTime: 5,
					MaxConnLifetime: uint64(30 * time.Minute),
				},
			},
			wantErr: true,
		},
		{
			name: "Empty configuration",
			args: args{
				cfg: &settings.SQLSetting{},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := database.NewConnection(tt.args.cfg)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewConnection() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && got == nil {
				t.Errorf("NewConnection() got nil connection for valid input")
			}
		})
	}
}
