package hasura_api

import (
	"github.com/go-resty/resty/v2"
	"github.com/gookit/config/v2"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type HasuraClient struct {
	Client      *resty.Client
	Config      *HasuraConfig
	adminSecret string
}

type HasuraClientOptions struct {
	configFilepaths []string
	envFilepaths    []string
}

type HasuraClientOption func(*HasuraClientOptions)

func WithConfigFilepath(filepath ...string) HasuraClientOption {
	return func(options *HasuraClientOptions) {
		options.configFilepaths = filepath
	}
}

func WithEnvFilepath(filepath ...string) HasuraClientOption {
	return func(options *HasuraClientOptions) {
		options.envFilepaths = filepath
	}
}

func NewHasuraClient(options ...HasuraClientOption) (*HasuraClient, error) {
	optionsStruct := new(HasuraClientOptions)

	for _, opt := range options {
		opt(optionsStruct)
	}

	if err := godotenv.Load(optionsStruct.envFilepaths...); err != nil {
		logrus.Warn("Error loading .env file", err)
	}

	conf, err := ConfigFromFile(optionsStruct.configFilepaths...)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// logrus.Info(conf.Endpoint)
	return &HasuraClient{
		Client:      resty.New(),
		adminSecret: config.Getenv("HASURA_GRAPHQL_ADMIN_SECRET", ""),
		Config:      conf,
	}, nil
}
