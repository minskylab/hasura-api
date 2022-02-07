package hasura_api

import (
	"time"

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
	Debug       bool
	adminSecret string
}

type MetadataClient struct {
	HasuraClient
}

type HasuraClientOptions struct {
	configFilepaths []string
	envFilepaths    []string
	debug           bool
	timeout         time.Duration
	literals        struct {
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

func WithDebug(debug bool) HasuraClientOption {
	return func(options *HasuraClientOptions) {
		options.debug = debug
	}
}

func WithTimeout(timeout time.Duration) HasuraClientOption {
	return func(options *HasuraClientOptions) {
		options.timeout = timeout
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
	opts := new(HasuraClientOptions)

	for _, opt := range options {
		opt(opts)
	}

	if opts.literals.endpoint != "" && opts.literals.adminSecret != "" {
		return newHasuraClientWithLiterals(opts.literals.endpoint, opts.literals.adminSecret)
	}

	if err := godotenv.Load(opts.envFilepaths...); err != nil {
		logrus.Warn("Error loading .env file", err)
	}

	conf, err := ConfigFromFile(opts.configFilepaths...)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	adminSecret := config.Getenv("HASURA_GRAPHQL_ADMIN_SECRET", "")

	client := &HasuraClient{
		client:      resty.New(),
		Debug:       opts.debug,
		adminSecret: adminSecret,
		Config:      conf,
	}

	client.client.SetTimeout(10 * time.Minute)

	if opts.timeout != time.Duration(0) {
		client.client.SetTimeout(opts.timeout)
	}

	client.Metadata = &MetadataClient{
		HasuraClient: *client,
	}

	return client, nil
}
