package hasura_api

import (
	"fmt"

	"github.com/minskylab/hasura-api/metadata"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func (r *MetadataClient) genericHasuraQuery(body metadata.MetadataQuery) (metadata.MetadataResponse, error) {
	endpoint := fmt.Sprintf("%s/v1/metadata", r.Config.Endpoint)

	logrus.Debug("requesting to hasura metadata query: ", body.Type)

	res, err := r.client.R().
		SetHeaders(map[string]string{
			"Content-Type":          "application/json",
			"X-Hasura-Role":         "admin",
			"X-Hasura-Admin-Secret": r.adminSecret,
		}).
		SetBody(body).
		Post(endpoint)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	switch res.StatusCode() {
	case 200, 201:
		switch successRes := res.Result().(type) {
		case metadata.SuccessResponse:
			return successRes, nil
		default:
			return nil, errors.Errorf("unexpected success response: %v", res.Result())
		}
	case 400:
		switch badRequestRes := res.Result().(type) {
		case metadata.BadRequestResponse:
			return badRequestRes, nil
		default:
			return nil, errors.Errorf("unexpected bad request response: %v", res.Result())
		}
	case 401:
		switch unauthorizedRes := res.Result().(type) {
		case metadata.UnauthorizedResponse:
			return unauthorizedRes, nil
		default:
			return nil, errors.Errorf("unexpected unauthorized response: %v", res.Result())
		}
	case 500:
		switch internalServerErrorRes := res.Result().(type) {
		case metadata.InternalServerErrorResponse:
			return internalServerErrorRes, nil
		default:
			return nil, errors.Errorf("unexpected internal server error response: %v", res.Result())
		}
	default:
		logrus.Warnf("unexpected response (code: %d): %v", res.StatusCode(), res.Result())
		return res.Result().(metadata.MetadataResponse), nil
	}
}
