package db

import (
	"fmt"
	"strings"
	"time"
)

const (
	defaultPoolMaxConns = 10
	defaultPoolMinConns = 1
)

type Config struct {
	Host                  string        `yaml:"host"`
	Port                  string        `yaml:"port"`
	User                  string        `yaml:"user"`
	Password              string        `yaml:"password"`
	DBName                string        `yaml:"db_name"`
	ApplicationName       string        `yaml:"application_name"`
	NoSQL                 bool          `yaml:"no_sql"`
	PoolMaxConns          int           `yaml:"pool_max_conns"`
	PoolMinConns          int           `yaml:"pool_min_conns"`
	PoolMaxConnLifetime   time.Duration `yaml:"pool_max_conn_lifetime"`
	PoolMaxConnIdleTime   time.Duration `yaml:"pool_max_conn_idle_time"`
	PoolHealthCheckPeriod time.Duration `yaml:"pool_health_check_period"`
}

func (c *Config) withDefaults() (conf Config) {
	if c != nil {
		conf = *c
	}

	if conf.PoolMaxConns == 0 {
		conf.PoolMaxConns = defaultPoolMaxConns
	}

	if conf.PoolMinConns == 0 {
		conf.PoolMinConns = defaultPoolMinConns
	}

	return
}

func (c Config) DSN() string {
	var builder strings.Builder

	if len(c.User) != 0 {
		builder.WriteString(fmt.Sprintf("user=%s ", c.User))
	}

	if len(c.Password) != 0 {
		builder.WriteString(fmt.Sprintf("password=%s ", c.Password))
	}

	if len(c.Host) != 0 {
		builder.WriteString(fmt.Sprintf("host=%s ", c.Host))
	}

	if len(c.Port) != 0 {
		builder.WriteString(fmt.Sprintf("port=%s ", c.Port))
	}

	if len(c.DBName) != 0 {
		builder.WriteString(fmt.Sprintf("dbname=%s ", c.DBName))
	}

	builder.WriteString(fmt.Sprintf("pool_max_conns=%d ", c.PoolMaxConns))
	builder.WriteString(fmt.Sprintf("pool_min_conns=%d ", c.PoolMinConns))

	if c.PoolMaxConnLifetime != 0 {
		builder.WriteString(fmt.Sprintf("pool_max_conn_lifetime=%s ", c.PoolMaxConnLifetime))
	}

	if c.PoolMaxConnIdleTime != 0 {
		builder.WriteString(fmt.Sprintf("pool_max_conn_idle_time=%s ", c.PoolMaxConnIdleTime))
	}

	if c.PoolHealthCheckPeriod != 0 {
		builder.WriteString(fmt.Sprintf("pool_health_check_period=%s ", c.PoolHealthCheckPeriod))
	}

	return builder.String()
}
