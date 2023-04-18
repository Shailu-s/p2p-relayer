package config

import "fmt"

// SqliteDbConfig is a struct holding the SQLite database connection configuration
type SqliteDbConfig struct {
	FileName string
}

// PostgresDbConfig is a struct holding the Postgres database connection configuration
type PostgresDbConfig struct {
	URL      string // DataBase URL for connection
	Driver   string // DataBase driver
	Host     string // DataBase host
	Port     int64
	Ssl      string // DataBase sslmode
	Name     string // DataBase name
	User     string // DataBase's user
	Password string // User's password
}

// AsPostgresDbUrl returns the Postgres database connection URL
func (c *PostgresDbConfig) AsPostgresDbUrl() string {
	return fmt.Sprintf(c.URL, c.Host, c.Port, c.User, c.Name, c.Password, c.Ssl)
}

func (v *viperConfig) GetSqliteConfig() SqliteDbConfig {
	return SqliteDbConfig{
		FileName: v.GetString("sqlite_file_name"),
	}
}

// ReadDBConfig Reads storage params from config.json
func (v *viperConfig) ReadDBConfig() PostgresDbConfig {
	return PostgresDbConfig{
		URL:      v.GetString("storage.url"),
		Driver:   v.GetString("storage.driver"),
		Host:     v.GetString("storage.host"),
		Port:     v.GetInt64("storage.port"),
		Ssl:      v.GetString("storage.ssl_mode"),
		Name:     v.GetString("storage.db_name"),
		User:     v.GetString("storage.user"),
		Password: v.GetString("storage.password"),
	}
}
