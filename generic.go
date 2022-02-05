package hasura_api

import (
	"fmt"

	"github.com/minskylab/hasura-api/metadata"
	"github.com/pkg/errors"
)

func (r *MetadataClient) genericHasuraQuery(body metadata.MetadataQuery) (metadata.MetadataResponse, error) {
	endpoint := fmt.Sprintf("%s/v1/metadata", r.Config.Endpoint)

	res, err := r.client.R().
		SetHeaders(map[string]string{
			"Content-Type":          "application/json",
			"X-Hasura-Role":         "admin",
			"X-Hasura-Admin-Secret": r.adminSecret,
		}).
		SetBody(body).
		Post(endpoint)
	if err != nil {
		// logrus.Warn(errors.WithStack(err))
		// logrus.Warn("response: ", res.StatusCode(), " ", res.String())
		// return nil
		return nil, errors.WithStack(err)
	}

	// logrus.Debug("response: ", res.StatusCode(), " ", res.String())
	switch res.StatusCode() {
	case 200:
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
		return nil, errors.Errorf("unexpected response: %v", res.Result())
	}
}

// func (r *HasuraClient) genericHasuraBulkQuery(bodies []metadata.MetadataQuery) error {
// 	endpoint := fmt.Sprintf("%s/v1/metadata", r.Config.Endpoint)

// 	body := metadata.MetadataQuery{
// 		Type: "bulk",
// 		Args: bodies,
// 	}

// 	res, err := r.Client.R().
// 		SetHeaders(map[string]string{
// 			"Content-Type":          "application/json",
// 			"X-Hasura-Role":         "admin",
// 			"X-Hasura-Admin-Secret": r.adminSecret,
// 		}).
// 		SetBody(body).
// 		Post(endpoint)
// 	if err != nil {
// 		logrus.Warn(errors.WithStack(err))
// 		logrus.Warn("response: ", res.StatusCode(), " ", res.String())
// 		return nil
// 	}

// 	logrus.Debug("response: ", res.StatusCode(), " ", res.String())

// 	return nil
// }
