package migration

import "crypto-dashboard-backend/configs/development"

type ConfigMigrate struct {
	DBSource     string
	MigrationURL string
}

func DefaultConfigMigrate() *ConfigMigrate {
	return &ConfigMigrate{
		DBSource:     development.GetEnvStr("DB_SOURCE", ""),
		MigrationURL: development.GetEnvStr("MIGRATION_URL", ""),
	}
}
