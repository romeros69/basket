package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

const pathToConfig = "./config/config.yml"

type (
	Config struct {
		App   `yaml:"app"`
		HTTP  `yaml:"http"`
		Log   `yaml:"logger"`
		Mongo `yaml:"mongo"`
		Neo4j `yaml:"neo4j"`
	}

	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	HTTP struct {
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}

	Mongo struct {
		MongoURL string `env-required:"true" yaml:"mongo_url" env:"MONGO_URL"`
		MongoDB  string `env-required:"true" yaml:"mongo_db" env:"MONGO_DB"`
	}

	Neo4j struct {
		Neo4jURL string `env-required:"true" yaml:"neo4j_url" env:"NEO4J_URL"`
		Neo4jLogin string `env-required:"true" yaml:"neo4j_login" env:"NEO4J_LOGIN"`
		Neo4jPassword string `env-required:"true" yaml:"neo4j_password" env:"NEO4J_PASSWORD"`
	}

	Log struct {
		Level string `env-required:"true" yaml:"log_level"   env:"LOG_LEVEL"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig(pathToConfig, cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
