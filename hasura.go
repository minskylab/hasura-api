package hasura_api

import (
	"github.com/go-resty/resty/v2"
	"github.com/gookit/config/v2"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type HasuraClient struct {
	Metadata    *MetadataClient
	client      *resty.Client
	Config      *HasuraConfig
	adminSecret string
}

type MetadataClient struct {
	HasuraClient
}

type HasuraClientOptions struct {
	configFilepaths []string
	envFilepaths    []string

	literals struct {
		endpoint    string
		adminSecret string
	}
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

func WithLiterals(hasuraEndpoint, hasuraAdminSecret string) HasuraClientOption {
	return func(options *HasuraClientOptions) {
		options.literals.endpoint = hasuraEndpoint
		options.literals.adminSecret = hasuraAdminSecret
	}
}

func newHasuraClientWithLiterals(hasuraEndpoint, hasuraAdminSecret string) (*HasuraClient, error) {
	client := &HasuraClient{
		client:      resty.New(),
		adminSecret: hasuraAdminSecret,
		Config: &HasuraConfig{
			Endpoint: hasuraEndpoint,
		},
	}

	client.Metadata = &MetadataClient{
		HasuraClient: *client,
	}

	return client, nil
}

func NewHasuraClient(options ...HasuraClientOption) (*HasuraClient, error) {
	optionsStruct := new(HasuraClientOptions)

	for _, opt := range options {
		opt(optionsStruct)
	}

	if optionsStruct.literals.endpoint != "" && optionsStruct.literals.adminSecret != "" {
		return newHasuraClientWithLiterals(optionsStruct.literals.endpoint, optionsStruct.literals.adminSecret)
	}

	if err := godotenv.Load(optionsStruct.envFilepaths...); err != nil {
		logrus.Warn("Error loading .env file", err)
	}

	conf, err := ConfigFromFile(optionsStruct.configFilepaths...)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	adminSecret := config.Getenv("HASURA_GRAPHQL_ADMIN_SECRET", "")

	client := &HasuraClient{
		client:      resty.New(),
		adminSecret: adminSecret,
		Config:      conf,
	}

	client.Metadata = &MetadataClient{
		HasuraClient: *client,
	}

	return client, nil
}
