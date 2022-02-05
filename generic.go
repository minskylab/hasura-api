package hasura_api

import (
	"fmt"

	"github.com/minskylab/hasura-api/metadata"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func (r *HasuraClient) genericHasuraQuery(body metadata.MetadataQuery) error {
	endpoint := fmt.Sprintf("%s/v1/metadata", r.Config.Endpoint)

	res, err := r.Client.R().
		SetHeaders(map[string]string{
			"Content-Type":          "application/json",
			"X-Hasura-Role":         "admin",
			"X-Hasura-Admin-Secret": r.adminSecret,
		}).
		SetBody(body).
		Post(endpoint)
	if err != nil {
		logrus.Warn(errors.WithStack(err))
		logrus.Warn("response: ", res.StatusCode(), " ", res.String())
		return nil
	}

	logrus.Debug("response: ", res.StatusCode(), " ", res.String())

	return nil
}

func (r *HasuraClient) genericHasuraBulkQuery(bodies []metadata.MetadataQuery) error {
	endpoint := fmt.Sprintf("%s/v1/metadata", r.Config.Endpoint)

	body := metadata.MetadataQuery{
		Type: "bulk",
		Args: bodies,
	}

	res, err := r.Client.R().
		SetHeaders(map[string]string{
			"Content-Type":          "application/json",
			"X-Hasura-Role":         "admin",
			"X-Hasura-Admin-Secret": r.adminSecret,
		}).
		SetBody(body).
		Post(endpoint)
	if err != nil {
		logrus.Warn(errors.WithStack(err))
		logrus.Warn("response: ", res.StatusCode(), " ", res.String())
		return nil
	}

	logrus.Debug("response: ", res.StatusCode(), " ", res.String())

	return nil
}
