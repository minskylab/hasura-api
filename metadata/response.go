package metadata

type MetadataResponse interface {
	GetResponse() interface{}
}

type SuccessResponse map[string]interface{}

func (r SuccessResponse) GetResponse() interface{} {
	return r
}

type BadRequestResponse struct {
	Path  string `json:"path"`
	Error string `json:"error"`
}

func (r BadRequestResponse) GetResponse() interface{} {
	return r
}

type UnauthorizedResponse struct {
	Error string `json:"error"`
}

func (r UnauthorizedResponse) GetResponse() interface{} {
	return r
}

type InternalServerErrorResponse struct {
	Error string `json:"error"`
}

func (r InternalServerErrorResponse) GetResponse() interface{} {
	return r
}
