package config

type Config struct {
	Port     string   `envconfig:"PORT" default:"8080"`
	DbConfig DbConfig `yaml:"db_config"`
}

type DbConfig struct {
	TableName string `yaml:"table_name" envconfig:"DB_TABLE_NAME"`
	Region    string `yaml:"region" envconfig:"DB_REGION"`
	Endpoint  string `yaml:"endpoint" envconfig:"DB_ENDPOINT"`
	AccessKey string `yaml:"access_key" envconfig:"DB_ACCESS_KEY"`
	SecretKey string `yaml:"secret_key" envconfig:"DB_SECRET_KEY"`
}
