package hasura_api

import (
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type HasuraConfig struct {
	Version           int     `yaml:"version"`
	Endpoint          string  `yaml:"endpoint"`
	MetadataDirectory string  `yaml:"metadata_directory"`
	Actions           Actions `yaml:"actions"`
}

type Actions struct {
	Kind                  string `yaml:"kind"`
	HandlerWebhookBaseurl string `yaml:"handler_webhook_baseurl"`
}

func defaultConfig() *HasuraConfig {
	return &HasuraConfig{
		Version:           3,
		Endpoint:          config.Getenv("HASURA_GRAPHQL_ENDPOINT", "http://localhost:8080"),
		MetadataDirectory: config.Getenv("HASURA_METADATA_DIRECTORY", "metadata"),
		Actions: Actions{
			Kind:                  config.Getenv("HASURA_ACTIONS_KIND", "synchronous"),
			HandlerWebhookBaseurl: config.Getenv("HASURA_ACTIONS_HANDLER_WEBHOOK_BASEURL", "http://localhost:3000"),
		},
	}
}

func ConfigFromFile(filepath ...string) (*HasuraConfig, error) {
	config.WithOptions(config.ParseEnv)
	config.AddDriver(yaml.Driver)
	config.WithOptions(func(opt *config.Options) {
		opt.TagName = "yaml"
	})

	if len(filepath) < 1 {
		return defaultConfig(), nil
	}

	err := config.LoadFiles(filepath...)
	if err == nil {
		c := new(HasuraConfig)
		return c, config.BindStruct("", c)
	}

	if err != nil {
		logrus.Error(errors.WithStack(err))
	}

	return defaultConfig(), nil
}

func MustConfigFromFile(filepath string) *HasuraConfig {
	c, err := ConfigFromFile(filepath)
	if err != nil {
		panic(err)
	}

	return c
}
