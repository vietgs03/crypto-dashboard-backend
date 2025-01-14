package database

import (
	"testing"
)

func TestNewConnection(t *testing.T) {
	type args struct {
		cfg *Config
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Valid configuration",
			args: args{
				cfg: &Config{
					Host:     "localhost",
					Port:     5432,
					User:     "postgres",
					Password: "password",
					DBName:   "testdb",
					SSLMode:  "disable",
					MaxConns: 10,
					Timeout:  5,
				},
			},
			wantErr: false,
		},
		{
			name: "Invalid host",
			args: args{
				cfg: &Config{
					Host:     "invalid_host",
					Port:     5432,
					User:     "postgres",
					Password: "password",
					DBName:   "testdb",
					SSLMode:  "disable",
					MaxConns: 10,
					Timeout:  5,
				},
			},
			wantErr: true,
		},
		{
			name: "Empty configuration",
			args: args{
				cfg: &Config{},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewConnection(tt.args.cfg)
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
